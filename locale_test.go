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

	tag, err = language.Parse("jam")
	ta.NoError(err)
	ta.Equal("jam", tag.String())
	l, err = localizer.GetLocaleData(tag)
	ta.NoError(err)
	ta.Equal(language.English, language.MustParse(l.Locale))

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
	ta.Equal(language.LatinAmericanSpanish, language.MustParse(l.Locale))
}
