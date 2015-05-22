package nl_AW

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "nl_AW",
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
