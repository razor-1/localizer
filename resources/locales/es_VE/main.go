package es_VE

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "es_VE",
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
