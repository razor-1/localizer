package luy

import "github.com/theplant/i18n/cldr"

var (
	symbols = cldr.Symbols{}
	formats = cldr.NumberFormats{Decimal: "", Currency: "¤#,##0.00;¤-\u00a0#,##0.00", Percent: ""}
)
