package en_NZ

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "en_NZ",
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
