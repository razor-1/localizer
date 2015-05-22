package en_SX

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "en_SX",
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
