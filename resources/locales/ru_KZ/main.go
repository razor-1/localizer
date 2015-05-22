package ru_KZ

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "ru_KZ",
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
