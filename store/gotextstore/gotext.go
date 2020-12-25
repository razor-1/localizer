package gotextstore

// small helper package which conforms a gotext locale to the localizer/store/TranslationStore interface
// see https://github.com/leonelquinteros/gotext/pull/47 for more information

import (
	"fmt"
	"github.com/leonelquinteros/gotext"
	"github.com/razor-1/localizer"
	"github.com/razor-1/localizer/store"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
)

type gotextStore struct {
	gotextLocale *gotext.Locale
}

//GetTranslationStore creates a store.TranslationStore from a gotext Locale
func GetTranslationStore(gotextLocale *gotext.Locale) store.TranslationStore {
	gts := &gotextStore{
		gotextLocale: gotextLocale,
	}

	return gts
}

func (gtstore *gotextStore) GetTranslations(tag language.Tag) (store.LocaleCatalog, error) {
	lc := store.NewLocaleCatalog(tag)
	lcData, err := localizer.GetLocaleData(tag)
	if err != nil {
		return lc, err
	}
	if lcData == nil {
		return lc, fmt.Errorf("unable to get locale data for tag %v", tag)
	}

	gotextTranslations := gtstore.gotextLocale.GetTranslations()

	storeTranslations := make(map[string]*store.Translation, len(gotextTranslations))
	for messageID, msg := range gotextTranslations {
		//populate the store.Translation with basic info
		st := &store.Translation{
			ID:       msg.ID,
			PluralID: msg.PluralID,
			String:   msg.Get(),
		}

		//convert gettext plurals into an explicit map of the plural form to the translation
		if msg.PluralID != "" && len(lcData.Plural.Cardinal.Forms) > 0 {
			plForms := make(map[plural.Form]string, len(lcData.Plural.Cardinal.Forms))
			for i, form := range lcData.Plural.Cardinal.Forms {
				plForms[form] = msg.GetN(i)
			}
			st.Plurals = plForms
		}
		storeTranslations[messageID] = st
		if msg.PluralID != "" {
			storeTranslations[msg.PluralID] = st
		}
	}

	lc.Translations = storeTranslations
	return lc, nil
}
