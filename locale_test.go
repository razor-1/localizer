package localizer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"

	"github.com/razor-1/localizer"
	"github.com/razor-1/localizer/store"
)

const (
	msgID         = "test"
	msgStr        = "testXlate"
	pluralMinutes = "%d minutes"
	pluralMinute  = "%d minute"
)

var enTag = language.Make("en")

func TestNewLocale(t *testing.T) {
	ta := assert.New(t)

	l, err := localizer.NewLocale(enTag)
	ta.NoError(err)
	ta.Equal(enTag, l.Tag)
	ta.NotZero(l.Calendar)
	ta.NotZero(l.Number)
	ta.NotZero(l.Plural)

	ta.NotNil(localizer.GetLocale(enTag))
}

func TestNamedParameters(t *testing.T) {
	ta := assert.New(t)

	bob := localizer.FmtParams{"name": "bob"}
	tests := []struct {
		format string
		params localizer.FmtParams
		out    string
	}{
		{
			format: "%(name)s is cool",
			params: bob,
			out:    "bob is cool",
		},
		{
			format: "what about %(name)",
			params: bob,
			out:    "what about bob",
		},
		{
			format: "%(num)d",
			params: localizer.FmtParams{"num": 22},
			out:    "22",
		},
		{
			format: "%(name)s has %(num)d and %(name) needs %(num2)d",
			params: localizer.FmtParams{"name": "bob", "num": 12, "num2": 20},
			out:    "bob has 12 and bob needs 20",
		},
		{
			format: "%(name)s has %(num)d and %(name)s needs %(num)d",
			params: localizer.FmtParams{"name": "bob", "num": 12},
			out:    "bob has 12 and bob needs 12",
		},
	}

	for _, test := range tests {
		t.Run(test.format, func(t *testing.T) {
			out := localizer.NamedParameters(test.format, test.params)
			ta.Equal(test.out, out)
		})
	}
}

type loader struct {
	catalog store.LocaleCatalog
}

func (ld *loader) GetTranslations(_ language.Tag) (store.LocaleCatalog, error) {
	return ld.catalog, nil
}

func getTestStore(t *testing.T) *loader {
	t.Helper()

	translations := make(map[string]*store.Translation)
	translations[msgID] = &store.Translation{
		ID:     msgID,
		String: msgStr,
	}

	pluralTrans := &store.Translation{
		ID:       "plural-minutes",
		PluralID: pluralMinutes,
		String:   pluralMinute,
		Plurals:  map[plural.Form]string{plural.One: pluralMinute, plural.Other: pluralMinutes},
	}
	translations[pluralTrans.PluralID] = pluralTrans

	testStore := &loader{
		catalog: store.LocaleCatalog{
			Tag:          enTag,
			Translations: translations,
		},
	}

	return testStore
}

func TestLocale_Load(t *testing.T) {
	ta := assert.New(t)

	l, err := localizer.NewLocale(enTag)
	ta.NoError(err)

	testStore := getTestStore(t)
	err = l.Load(testStore)
	ta.NoError(err)
	ta.Equal(msgStr, l.Get(msgID))

	prt := l.NewPrinter()
	ta.Equal(msgStr, prt.Sprintf(msgID))

	ta.Equal("1 minute", l.GetPlural(pluralMinutes, 1, 1))
	ta.Equal("2 minutes", l.GetPlural(pluralMinutes, 2, 2))
}

func TestNewLocaleWithStore(t *testing.T) {
	ta := assert.New(t)
	testStore := getTestStore(t)

	l, err := localizer.NewLocaleWithStore(enTag, testStore)
	ta.NoError(err)
	ta.NotNil(l)
	ta.Equal(enTag, l.Tag)
}

func TestLocale_GetTranslations(t *testing.T) {
	ta := assert.New(t)
	testStore := getTestStore(t)
	l, err := localizer.NewLocaleWithStore(enTag, testStore)
	ta.NoError(err)

	cat := l.GetTranslations()
	ta.Len(cat, len(testStore.catalog.Translations))
}

func TestGetLocaleData(t *testing.T) {
	ta := assert.New(t)

	tag, err := language.Parse("ase")
	ta.NoError(err)
	ta.Equal("ase", tag.String())
	l, err := localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.AmericanEnglish, language.MustParse(l.Locale))

	tag, err = language.Parse("asf")
	ta.NoError(err)
	ta.Equal("asf", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal("en-AU", language.MustParse(l.Locale).String())

	tag, err = language.Parse("jam")
	ta.NoError(err)
	ta.Equal("jam", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.English, language.MustParse(l.Locale))

	tag, err = language.Parse("ht")
	ta.NoError(err)
	ta.Equal("ht", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal("ht", language.MustParse(l.Locale).String())

	tag, err = language.Parse("ca-valencia")
	ta.NoError(err)
	ta.Equal("ca-valencia", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.Catalan, language.MustParse(l.Locale))

	tag, err = language.Parse("quc")
	ta.NoError(err)
	ta.Equal("quc", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal("quc", language.MustParse(l.Locale).String())

	tag, err = language.Parse("lir")
	ta.NoError(err)
	ta.Equal("lir", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.English, language.MustParse(l.Locale))

	tag, err = language.Parse("bzj")
	ta.NoError(err)
	ta.Equal("bzj", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.English, language.MustParse(l.Locale))

	tag, err = language.Parse("fon")
	ta.NoError(err)
	ta.Equal("fon", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal("ee", language.MustParse(l.Locale).String())

	tag, err = language.Parse("tw")
	ta.NoError(err)
	ta.Equal("tw", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal("ak", language.MustParse(l.Locale).String())

	tag, err = language.Parse("quh")
	ta.NoError(err)
	ta.Equal("quh", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.LatinAmericanSpanish, language.MustParse(l.Locale))

	tag, err = language.Parse("psr")
	ta.NoError(err)
	ta.Equal("psr", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.EuropeanPortuguese, language.MustParse(l.Locale))

	tag, err = language.Parse("bzs")
	ta.NoError(err)
	ta.Equal("bzs", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.BrazilianPortuguese, language.MustParse(l.Locale))

	tag, err = language.Parse("tll")
	ta.NoError(err)
	ta.Equal("tll", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal("ln", language.MustParse(l.Locale).String())

	tag, err = language.Parse("gpe")
	ta.NoError(err)
	ta.Equal("gpe", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.English, language.MustParse(l.Locale))

	tag, err = language.Parse("gcf")
	ta.NoError(err)
	ta.Equal("gcf", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.French, language.MustParse(l.Locale))

	tag, err = language.Parse("rmn-BG")
	ta.NoError(err)
	ta.Equal("rmn-BG", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.Bulgarian, language.MustParse(l.Locale))

	tag, err = language.Parse("csq")
	ta.NoError(err)
	ta.Equal("csq", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.Croatian, language.MustParse(l.Locale))

	// ts-MZ uses Portuguese for CLDR data but still has its own tag and translations
	// there is a special override to force this, so that we don't get Tsonga CLDR data
	tag, err = language.Parse("ts-MZ")
	ta.NoError(err)
	ta.Equal("ts-MZ", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.EuropeanPortuguese, language.MustParse(l.Locale))
	loc, err := localizer.NewLocale(tag)
	ta.NoError(err)
	ta.Equal(tag, loc.Tag)

	tag, err = language.Parse("rmc-SK")
	ta.NoError(err)
	ta.Equal("rmc-SK", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.Slovak, language.MustParse(l.Locale))

	tag, err = language.Parse("rmy")
	ta.NoError(err)
	ta.Equal("rmy", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.Romanian, language.MustParse(l.Locale))

	tag, err = language.Parse("swc")
	ta.NoError(err)
	ta.Equal("swc", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.French, language.MustParse(l.Locale))

	tag, err = language.Parse("rms")
	ta.NoError(err)
	ta.Equal("rms", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.Romanian, language.MustParse(l.Locale))

	tag, err = language.Parse("gug")
	ta.NoError(err)
	ta.Equal("gug", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal("gn", language.MustParse(l.Locale).String())

	tag, err = language.Parse("swc-CD-x-katanga")
	ta.NoError(err)
	ta.Equal("swc-CD-x-katanga", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.French, language.MustParse(l.Locale))

	tag, err = language.Parse("rmn-Cyrl")
	ta.NoError(err)
	ta.Equal("rmn-Cyrl", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.Macedonian, language.MustParse(l.Locale))

	tag, err = language.Parse("sop")
	ta.NoError(err)
	ta.Equal("sop", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal("lu", language.MustParse(l.Locale).String())
}
