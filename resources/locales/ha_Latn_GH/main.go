package ha_Latn_GH

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "ha_Latn_GH",
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
