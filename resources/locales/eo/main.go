package eo

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "eo",
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
