package nl_BQ

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "nl_BQ",
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
