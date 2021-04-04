package localizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestLocale_DisplayName(t *testing.T) {
	ta := assert.New(t)

	l, err := NewLocale(language.English)
	ta.NoError(err)
	ta.Equal("Spanish", l.DisplayName(language.Spanish))
	ta.Equal("Brazilian Portuguese", l.DisplayName(language.BrazilianPortuguese))
	ta.Equal("European Portuguese", l.DisplayName(language.EuropeanPortuguese))

	l, err = NewLocale(language.Spanish)
	ta.NoError(err)
	ta.Equal("español", l.DisplayName(language.Spanish))
	ta.Equal("inglés", l.DisplayName(language.English))
	ta.Equal("inglés estadounidense", l.DisplayName(language.AmericanEnglish))
}

func ExampleLocale_DisplayName() {
	loc, _ := NewLocale(language.Spanish)
	fmt.Println(loc.DisplayName(language.English))
	//	Output: inglés
}
