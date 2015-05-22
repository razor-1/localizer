package bs_Cyrl

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "bs_Cyrl",
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
