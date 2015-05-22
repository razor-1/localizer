package en_PN

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "en_PN",
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
