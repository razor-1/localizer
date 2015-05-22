package so_DJ

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "so_DJ",
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
