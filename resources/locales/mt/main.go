package mt

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "mt",
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
