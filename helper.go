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

func (l *Locale) FmtDateFull(tim time.Time) (string, error) { return l.Calendar.FmtDateFull(tim) }
func (l *Locale) FmtDateLong(tim time.Time) (string, error) { return l.Calendar.FmtDateLong(tim) }
func (l *Locale) FmtDateMedium(tim time.Time) (string, error) {
	return l.Calendar.FmtDateMedium(tim)
}
func (l *Locale) FmtDateShort(tim time.Time) (string, error) { return l.Calendar.FmtDateShort(tim) }

func (l *Locale) FmtDateTimeFull(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeFull(tim)
}
func (l *Locale) FmtDateTimeLong(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeLong(tim)
}
func (l *Locale) FmtDateTimeMedium(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeMedium(tim)
}
func (l *Locale) FmtDateTimeShort(tim time.Time) (string, error) {
	return l.Calendar.FmtDateTimeShort(tim)
}
func (l *Locale) FmtTimeFull(tim time.Time) (string, error) { return l.Calendar.FmtTimeFull(tim) }
func (l *Locale) FmtTimeLong(tim time.Time) (string, error) { return l.Calendar.FmtTimeLong(tim) }
func (l *Locale) FmtTimeMedium(tim time.Time) (string, error) {
	return l.Calendar.FmtTimeMedium(tim)
}
func (l *Locale) FmtTimeShort(tim time.Time) (string, error) { return l.Calendar.FmtTimeShort(tim) }

func (l *Locale) FmtCurrency(currency string, number interface{}) (formatted string, err error) {
	return l.Number.FmtCurrency(currency, toFloat64(number))
}
func (l *Locale) FmtCurrencyWhole(currency string, number interface{}) (formatted string, err error) {
	return l.Number.FmtCurrencyWhole(currency, toFloat64(number))
}
func (l *Locale) FmtNumber(number interface{}) string {
	return l.Number.FmtNumber(toFloat64(number))
}
func (l *Locale) FmtNumberWhole(number interface{}) string {
	return l.Number.FmtNumberWhole(toFloat64(number))
}
func (l *Locale) FmtPercent(number interface{}) string {
	return l.Number.FmtPercent(toFloat64(number))
}
