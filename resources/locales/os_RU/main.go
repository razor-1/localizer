package os_RU

import "github.com/theplant/i18n/cldr"

var Locale = &cldr.Locale{
	Locale: "os_RU",
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
