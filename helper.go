package localizer

import (
	"reflect"
	"time"
)

var floatType = reflect.TypeOf(float64(0))

func toFloat64(unk interface{}) float64 {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0
	}
	fv := v.Convert(floatType)
	return fv.Float()
}

//FmtDateFull returns the CLDR Full form of the date, e.g. "Tuesday, September 14, 1999"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtDateFull(tim time.Time) (string, error) { return l.Calendar.FmtDateFull(tim) }

//FmtDateLong returns the CLDR Long form of the date, e.g. "September 14, 1999"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtDateLong(tim time.Time) (string, error) { return l.Calendar.FmtDateLong(tim) }

//FmtDateMedium returns the CLDR Medium form of the date, e.g. "Sep 14, 1999"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtDateMedium(tim time.Time) (string, error) { return l.Calendar.FmtDateMedium(tim) }

//FmtDateShort returns the CLDR Short form of the date, e.g. "9/14/99"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtDateShort(tim time.Time) (string, error) { return l.Calendar.FmtDateShort(tim) }

//FmtDateTimeFull returns the CLDR Full form of the date and time, e.g.
//"Tuesday, September 14, 1999 at 1:04:20 PM GMT+00:00"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtDateTimeFull(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeFull(tim)
}

//FmtDateTimeLong returns the CLDR Long form of the date and time, e.g. "September 14, 1999 at 1:04:20 PM GMT+00:00"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtDateTimeLong(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeLong(tim)
}

//FmtDateTimeMedium returns the CLDR Medium form of the date and time, e.g. "Sep 14, 1999, 1:04:20 PM"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtDateTimeMedium(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeMedium(tim)
}

//FmtDateTimeShort returns the CLDR Short form of the date and time, e.g. "9/14/99, 1:04 PM"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtDateTimeShort(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeShort(tim)
}

//FmtTimeFull returns the CLDR Full form of the time, e.g. "1:04:20 PM GMT+00:00"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtTimeFull(tim time.Time) (string, error) { return l.Calendar.FmtTimeFull(tim) }

//FmtTimeLong returns the CLDR Long form of the time, e.g. "1:04:20 PM GMT+00:00"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtTimeLong(tim time.Time) (string, error) { return l.Calendar.FmtTimeLong(tim) }

//FmtTimeMedium returns the CLDR Medium form of the time, e.g. "1:04:20 PM"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtTimeMedium(tim time.Time) (string, error) { return l.Calendar.FmtTimeMedium(tim) }

//FmtTimeShort returns the CLDR Short form of the time, e.g. "1:04 PM"
//see http://cldr.unicode.org/translation/date-time-1/date-time-patterns for more information
func (l *Locale) FmtTimeShort(tim time.Time) (string, error) { return l.Calendar.FmtTimeShort(tim) }

//FmtCurrency returns the number formatted as a currency for the locale, e.g. "$123.45"
func (l *Locale) FmtCurrency(currency string, number interface{}) (formatted string, err error) {
	return l.Number.FmtCurrency(currency, toFloat64(number))
}

//FmtCurrencyWhole returns the number formatted as a currency for the locale, excluding any decimal portion, e.g. "$123"
func (l *Locale) FmtCurrencyWhole(currency string, number interface{}) (formatted string, err error) {
	return l.Number.FmtCurrencyWhole(currency, toFloat64(number))
}

//FmtNumber returns the number formatted for the locale, e.g. "123.45" or "123,45"
func (l *Locale) FmtNumber(number interface{}) string { return l.Number.FmtNumber(toFloat64(number)) }

//FmtNumberWhole returns the number formatted for the locale, excluding any decimal portion, e.g. "123.45" or "123,45"
func (l *Locale) FmtNumberWhole(number interface{}) string {
	return l.Number.FmtNumberWhole(toFloat64(number))
}

//FmtPercent returns the number formatted as a percentage for the locale, e.g. 0.09 would return "9 %"
func (l *Locale) FmtPercent(number interface{}) string { return l.Number.FmtPercent(toFloat64(number)) }
