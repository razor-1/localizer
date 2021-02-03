package store

import "golang.org/x/text/language"

//LocaleCatalog represents the translations provided by a TranslationStore
type LocaleCatalog struct {
	//Tag is which locale was actually used to find these translations (e.g. I asked for en-US and am getting en)
	Tag language.Tag
	//Path is the filesystem path for the source of the catalog
	Path         string
	Translations map[string]*Translation
}

//NewLocaleCatalog creates an empty LocaleCatalog for the supplied language
func NewLocaleCatalog(tag language.Tag) LocaleCatalog {
	return LocaleCatalog{
		Tag:          tag,
		Path:         "",
		Translations: make(map[string]*Translation),
	}
}

//TranslationStore is implemented by packages that offer a way to load translations.
type TranslationStore interface {
	GetTranslations(tag language.Tag) (LocaleCatalog, error)
}
