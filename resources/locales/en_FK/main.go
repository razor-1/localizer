package en_FK

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "en_FK",
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
