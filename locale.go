package localizer

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/locales"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"

	"github.com/razor-1/localizer/store"
)

type Locale struct {
	Tag                language.Tag
	Number             cldr.Number
	Calendar           cldr.Calendar
	Plural             cldr.Plural
	catalog            *catalog.Builder
	translations       map[string]*store.Translation
	trMutex            sync.RWMutex
	pluralTranslations map[string]*store.Translation
	plMutex            sync.RWMutex
}

var llMutex sync.RWMutex
var loadedLocales = make(map[language.Tag]*Locale)

//NewLocale instantiates a new Locale for the supplied language tag, and adds the default domain from the global
//LocaleConfig.
func NewLocale(tag language.Tag) (loc *Locale, err error) {
	if loaded := GetLocale(tag); loaded != nil {
		return loaded, nil
	}
	localeData, err := GetLocaleData(tag)
	if err != nil {
		return
	}
	l := Locale{
		Tag:                tag,
		Number:             localeData.Number,
		Calendar:           localeData.Calendar,
		Plural:             localeData.Plural,
		catalog:            catalog.NewBuilder(),
		pluralTranslations: make(map[string]*store.Translation),
		translations:       make(map[string]*store.Translation),
	}

	llMutex.Lock()
	loadedLocales[tag] = &l
	llMutex.Unlock()

	return &l, nil
}

type FmtParams map[string]interface{}

var namedParameter = regexp.MustCompile(`%\((\w+?)\)(\S)?`)

//NamedParameters does string formatting on python-style format strings like "Hello %(name)s"
func NamedParameters(format string, params FmtParams) string {
	args := make([]string, len(params)*2)
	i := 0
	//if we have something like "%(name)" without a trailing format specifier, use defaultFormat
	const defaultFormat = "s"
	matches := namedParameter.FindAllStringSubmatch(format, -1)
	for _, match := range matches {
		param, ok := params[match[1]]
		if !ok {
			continue
		}
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

//GetLocaleData finds the best match for tag and returns a *cldr.Locale, which contains data populated from the
//Unicode CLDR.
func GetLocaleData(tag language.Tag) (*cldr.Locale, error) {
	//find the closest valid language for the supplied tag
	for {
		if loc, ok := locales.LocaleData[tag]; ok {
			return loc(), nil
		}
		tag = tag.Parent()
		if tag.IsRoot() {
			break
		}
	}

	return nil, errors.New("tag could not match a known locale")
}

func (l *Locale) Load(source store.TranslationStore) error {
	lc, err := source.GetTranslations(l.Tag)
	if err != nil {
		return err
	}

	l.plMutex.Lock()
	l.trMutex.Lock()
	defer l.plMutex.Unlock()
	defer l.trMutex.Unlock()
	for msgID, msg := range lc.Translations {
		if msg.PluralID != "" {
			l.pluralTranslations[msg.PluralID] = msg
		}
		_ = l.catalog.SetString(l.Tag, msgID, msg.String)
		l.translations[msgID] = msg
	}

	return nil
}

func (l *Locale) Get(key string) string {
	l.trMutex.RLock()
	defer l.trMutex.RUnlock()
	if tr, ok := l.translations[key]; ok {
		return tr.Get()
	}

	return key
}

//AddDomain adds a domain of translations to the locale and sets them in the catalog.
//func (l *Locale) AddDomain(domain string) error {
//	err := l.gotextLocale.AddDomain(domain)
//	if err != nil {
//		return err
//	}
//
//	for msgID, msg := range l.GetAllTranslations() {
//		err := l.catalog.SetString(l.Tag, msgID, msg)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

//NewPrinter creates a message.Printer for the Locale
func (l *Locale) NewPrinter() *message.Printer {
	return message.NewPrinter(l.Tag, message.Catalog(l.catalog))
}

//GetPlural determines which plural form should be used for the supplied number. number can be an int type (including
//int64) or a float formatted as a string. It then determines which gettext index number should be used for this
//plural form and retrieves it from the catalog.
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

// GetLocale returns a pointer to an exiting locale object (or nil if if doesn't exist)
func GetLocale(tag language.Tag) *Locale {
	llMutex.RLock()
	defer llMutex.RUnlock()
	return loadedLocales[tag]
}
