package ar_LB

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "ar_LB",
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
