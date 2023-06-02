package localizer

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"

	"github.com/razor-1/localizer/store"
)

// Locale is the struct providing access to most of the localizer features
type Locale struct {
	Tag            language.Tag
	Number         cldr.Number
	Calendar       cldr.Calendar
	Plural         cldr.Plural
	Languages      cldr.Languages
	Territories    cldr.Territories
	DisplayPattern cldr.LocaleDisplayPattern
	catalog        *catalog.Builder
	translations   map[string]*store.Translation
	trMutex        sync.RWMutex
}

var llMutex sync.RWMutex
var loadedLocales = make(map[language.Tag]*Locale)

// NewLocale instantiates a new Locale for the supplied language tag. It does not load any translations, so it's
// useful when you only need to use the CLDR data (number/calendar etc). Use NewLocaleWithStore to load translations
// at initialization time, or call Load on the Locale returned by this function.
func NewLocale(tag language.Tag) (loc *Locale, err error) {
	if loaded := GetLocale(tag); loaded != nil {
		return loaded, nil
	}
	localeData, err := GetLocaleData(tag)
	if err != nil {
		return
	}
	l := Locale{
		Tag:            tag,
		Number:         localeData.Number,
		Calendar:       localeData.Calendar,
		Plural:         localeData.Plural,
		Languages:      localeData.Languages,
		Territories:    localeData.Territories,
		DisplayPattern: localeData.Display,
		catalog:        catalog.NewBuilder(),
		translations:   make(map[string]*store.Translation),
	}

	llMutex.Lock()
	defer llMutex.Unlock()
	loadedLocales[tag] = &l

	return &l, nil
}

// NewLocaleWithStore instantiates a new Locale for the supplied language tag. It loads the translations from the source
// store.TranslationStore. Use this if you know you want your Locale to be populated with translations.
func NewLocaleWithStore(tag language.Tag, source store.TranslationStore) (loc *Locale, err error) {
	loc, err = NewLocale(tag)
	if err != nil {
		return
	}

	err = loc.Load(source)
	return
}

// GetLocale returns a pointer to an existing, already-loaded locale (or nil if if doesn't exist)
func GetLocale(tag language.Tag) *Locale {
	llMutex.RLock()
	defer llMutex.RUnlock()
	return loadedLocales[tag]
}

// FmtParams contains substitutions for python-style format strings
type FmtParams map[string]interface{}

var namedParameter = regexp.MustCompile(`%\((\w+?)\)(\S)?`)

// NamedParameters does string formatting on python-style format strings like "Hello %(name)s".
func NamedParameters(format string, params FmtParams) string {
	// if we have something like "%(name)" without a trailing format specifier, use defaultFormat
	const defaultFormat = "s"

	matches := namedParameter.FindAllStringSubmatch(format, -1)
	found := make(map[string]bool, len(matches))
	args := make([]string, len(matches)*2)
	i := 0

	for _, match := range matches {
		param, ok := params[match[1]]
		if !ok {
			continue
		}
		if _, ok := found[match[0]]; ok {
			// same name multiple times in string. replacer will replace all of them
			continue
		}
		found[match[0]] = true
		args[i] = match[0]
		fmtStr := "%"
		if len(match) < 3 || match[2] == "" {
			fmtStr += defaultFormat
		} else {
			fmtStr += match[2]
		}
		args[i+1] = fmt.Sprintf(fmtStr, param)
		i += 2
	}

	return strings.NewReplacer(args...).Replace(format)
}

func getFallbackTag(tag language.Tag) (language.Tag, error) {
	switch tag.String() {
	case "ase":
		// there is no CLDR locale for American Sign Language, so we need to fall back to en-US
		return language.AmericanEnglish, nil
	case "ht":
		// Haitian Creole is simply not part of CLDR, as of 2023-01-23. This is very surprising.
		return language.Make("fr-HT"), nil
	case "vec-BR":
		return language.Italian, nil
	}

	return language.Tag{}, fmt.Errorf("no fallback for tag %s", tag.String())
}

// GetLocaleData finds the best match for tag and returns a *cldr.Locale, which contains data populated from the
// Unicode CLDR.
func GetLocaleData(tag language.Tag) (*cldr.Locale, error) {
	// find the closest valid language for the supplied tag
	originalTag := tag
	for {
		if loc, err := resources.GetLocale(tag); err == nil {
			return loc, nil
		}
		tag = tag.Parent()
		if tag.IsRoot() {
			break
		}
	}

	// see if we have a fallback
	if fallbackTag, err := getFallbackTag(originalTag); err == nil {
		if loc, err := resources.GetLocale(fallbackTag); err == nil {
			return loc, nil
		}
	}

	return nil, fmt.Errorf("localizer.GetLocaleData: tag %s could not match a known locale", tag.String())
}

// Load retrieves all translations from the supplied store.TranslationStore and prepares them for use in this Locale.
func (l *Locale) Load(source store.TranslationStore) error {
	lc, err := source.GetTranslations(l.Tag)
	if err != nil {
		return err
	}

	l.trMutex.Lock()
	defer l.trMutex.Unlock()
	for msgID, msg := range lc.Translations {
		_ = l.catalog.SetString(l.Tag, msgID, msg.String)
		l.translations[msgID] = msg
	}

	return nil
}

// Get returns the raw translated string from the catalog, with no formatting.
func (l *Locale) Get(key string) string {
	l.trMutex.RLock()
	defer l.trMutex.RUnlock()
	if tr, ok := l.translations[key]; ok {
		return tr.Get()
	}

	return key
}

// GetTranslations returns the entire catalog of translations currently loaded for this locale. This allows for
// enumeration of the catalog. Note that it returns a copy so that the internal store can continue to be protected
// by mutexes.
func (l *Locale) GetTranslations() map[string]store.Translation {
	l.trMutex.RLock()
	defer l.trMutex.RUnlock()

	trans := make(map[string]store.Translation, len(l.translations))
	for k, v := range l.translations {
		trans[k] = *v
	}

	return trans
}

// CountTranslations returns the number of translations currently loaded in this locale.
func (l *Locale) CountTranslations() int {
	return len(l.translations)
}

// NewPrinter creates a message.Printer for the Locale
func (l *Locale) NewPrinter() *message.Printer {
	return message.NewPrinter(l.Tag, message.Catalog(l.catalog))
}

// GetPlural determines which plural form should be used for the supplied number. number can be an int type (including
// int64) or a float formatted as a string. It then determines which plural form should be used by calling the cldr
// plural function, and returns the corresponding plural translation, formatted with the optional vars.
func (l *Locale) GetPlural(pluralID string, number interface{}, vars ...interface{}) string {
	ops, err := cldr.NewOperands(number)
	if err != nil {
		return pluralID
	}

	form := l.Plural.Cardinal.Func(ops)
	var msg string
	l.trMutex.RLock()
	defer l.trMutex.RUnlock()
	if tr, ok := l.translations[pluralID]; ok {
		msg = tr.GetPlural(form)
	} else {
		msg = pluralID
	}

	if len(vars) > 0 {
		return fmt.Sprintf(msg, vars...)
	}
	return msg
}
