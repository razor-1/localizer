package de_LI

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "de_LI",
	Number: cldr.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar: calendar,
}

func init() {
	cldr.RegisterLocale(Locale)
}
