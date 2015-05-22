package cldr

type PluralRule string

const (
	PluralRuleZero  PluralRule = "zero"  // zero
	PluralRuleOne              = "one"   // singular
	PluralRuleTwo              = "two"   // dual
	PluralRuleFew              = "few"   // paucal
	PluralRuleMany             = "many"  // also used for fractions if they have a separate class
	PluralRuleOther            = "other" // required—general plural form—also used if the language only has a single form
)

// Number should be one of these types:
//  int, float
type NumberValue interface{}

type PluralRuler interface {
	FindRule(count NumberValue) PluralRule
}

var pluralRules = map[string]PluralRuler{}

func RegisterPluralRule(locale string, ruler PluralRuler) {
	pluralRules[locale] = ruler
}

func FindRule(locale string, count NumberValue) (rule PluralRule) {
	ruler, ok := pluralRules[locale]
	if !ok {
		return PluralRuleOther
	}
	return ruler.FindRule(count)
}

type PluralRulerFunc func(count NumberValue) PluralRule

func (p PluralRulerFunc) FindRule(count NumberValue) PluralRule {
	return p(count)
}

func init() {
	RegisterPluralRule("en", PluralRulerFunc(func(count NumberValue) PluralRule {
		switch count.(type) {
		case int, int32, int64, uint, uint32, uint64:
			if count == 1 {
				return PluralRuleOne
			}
		}
		return PluralRuleOther
	}))
}
