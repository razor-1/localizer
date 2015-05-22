package se_SE

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "se_SE",
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
