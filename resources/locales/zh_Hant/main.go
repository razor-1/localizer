package zh_Hant

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "zh_Hant",
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
