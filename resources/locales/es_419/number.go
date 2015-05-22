package es_419

import "github.com/theplant/i18n/cldr"

var (
	symbols = cldr.Symbols{Decimal: ".", Group: ",", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "0", Currency: "¤#,##0.00", Percent: ""}
)
