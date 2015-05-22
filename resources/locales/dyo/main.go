package dyo

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "dyo",
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
