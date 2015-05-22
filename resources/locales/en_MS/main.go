package en_MS

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "en_MS",
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
