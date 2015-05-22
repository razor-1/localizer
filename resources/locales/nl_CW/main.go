package nl_CW

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "nl_CW",
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
