package es_NI

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "es_NI",
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
