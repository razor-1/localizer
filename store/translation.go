package store

import "golang.org/x/text/feature/plural"

// Translation is the struct for the Translations parsed via Po or Mo files and all coming parsers
type Translation struct {
	ID, PluralID, String string
	Plurals              map[plural.Form]string
}

// NewTranslation returns the initialized Translation
func NewTranslation() *Translation {
	tr := new(Translation)
	tr.Plurals = make(map[plural.Form]string, 2)

	return tr
}

// Get returns the string of the translation
func (t *Translation) Get() string {
	if t.String != "" {
		return t.String
	}

	// Return untranslated id by default
	return t.ID
}

//GetPlural returns the requested plural form, if available
func (t *Translation) GetPlural(form plural.Form) string {
	if tr, ok := t.Plurals[form]; ok {
		return tr
	}

	return t.PluralID
}
