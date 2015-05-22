package es_PH

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "es_PH",
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
