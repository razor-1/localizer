package en_NG

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "en_NG",
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
