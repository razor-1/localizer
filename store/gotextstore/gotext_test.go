package gotextstore

import (
	"testing"

	"github.com/leonelquinteros/gotext"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
)

const (
	enUS           = "enUS"
	domainMessages = "messages"
)

func TestGotextStore_GetTranslations(t *testing.T) {
	gotextLocale := gotext.NewLocale("testdata/", enUS)
	gotextLocale.AddDomain(domainMessages)

	ta := assert.New(t)
	tag, err := language.Parse(enUS)
	ta.NoError(err)
	ta.Equal(language.AmericanEnglish, tag)

	gts := GetTranslationStore(gotextLocale)
	lc, err := gts.GetTranslations(tag)
	ta.NoError(err)

	const msgIDLanguage = "language"
	lang, ok := lc.Translations[msgIDLanguage]
	ta.True(ok)
	ta.Equal(enUS, lang.Get())
	ta.Equal(msgIDLanguage, lang.ID)
	ta.Nil(lang.Plurals)

	const msgIDPlural = "One with var: %s"
	const pluralMsgID = "Several with vars: %s"
	pl, ok := lc.Translations[msgIDPlural]
	ta.True(ok)
	ta.Equal(pluralMsgID, pl.PluralID)
	ta.Equal("This one is the singular: %s", pl.GetPlural(plural.One))
	ta.Equal("This one is the plural: %s", pl.GetPlural(plural.Other))

	plid, ok := lc.Translations[pluralMsgID]
	ta.True(ok)
	ta.Equal(pl, plid)
}
