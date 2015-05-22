package sq_MK

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "sq_MK",
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
