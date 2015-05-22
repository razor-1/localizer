package ar_MR

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "ar_MR",
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
