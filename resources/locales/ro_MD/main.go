package ro_MD

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "ro_MD",
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
