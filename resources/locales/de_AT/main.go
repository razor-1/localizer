package de_AT

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "de_AT",
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
