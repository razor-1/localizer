package fa

import "github.com/theplant/i18n/cldr"

var (
	symbols = cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u200e−", Percent: "٪", PerMille: "؉"}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "\u200e¤#,##0.00", Percent: "#,##0%"}
)
