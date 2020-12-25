package gotextstore

import (
	"github.com/leonelquinteros/gotext"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"testing"
)

const (
	en_US          = "en_US"
	domainMessages = "messages"
)

func TestGotextStore_GetTranslations(t *testing.T) {
	gotextLocale := gotext.NewLocale("testdata/", en_US)
	gotextLocale.AddDomain(domainMessages)

	ta := assert.New(t)
	tag, err := language.Parse(en_US)
	ta.NoError(err)
	ta.Equal(language.AmericanEnglish, tag)

	gts := GetTranslationStore(gotextLocale)
	lc, err := gts.GetTranslations(tag)
	ta.NoError(err)

	const msgIdLanguage = "language"
	lang, ok := lc.Translations[msgIdLanguage]
	ta.True(ok)
	ta.Equal(en_US, lang.Get())
	ta.Equal(msgIdLanguage, lang.ID)
	ta.Nil(lang.Plurals)

	const msgIdPlural = "One with var: %s"
	const pluralMsgId = "Several with vars: %s"
	pl, ok := lc.Translations[msgIdPlural]
	ta.True(ok)
	ta.Equal(pluralMsgId, pl.PluralID)
	ta.Equal("This one is the singular: %s", pl.GetPlural(plural.One))
	ta.Equal("This one is the plural: %s", pl.GetPlural(plural.Other))

	plid, ok := lc.Translations[pluralMsgId]
	ta.True(ok)
	ta.Equal(pl, plid)
}
