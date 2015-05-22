package ff_GN

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "ff_GN",
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
