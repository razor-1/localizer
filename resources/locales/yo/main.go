package yo

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "yo",
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
