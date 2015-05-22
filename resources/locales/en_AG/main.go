package en_AG

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "en_AG",
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
