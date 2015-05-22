package br

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "br",
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
