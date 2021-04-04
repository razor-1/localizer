package localizer

import (
	"strings"

	"golang.org/x/text/language"
)

const (
	//{0} and {1} are the tokens to be replaced in the cldr locale display patterns
	displayToken0 = "{0}"
	displayToken1 = "{1}"
)

//DisplayName returns the formatted display name for the provided language tag in the locale.
func (l *Locale) DisplayName(tag language.Tag) string {

	base, script, region := tag.Raw()
	var haveScript, haveRegion bool
	lang, ok := l.Languages[base.String()]
	if !ok {
		return base.String()
	}

	if region.String() != "ZZ" {
		haveRegion = true
	}
	if script.String() != "Zzzz" {
		haveScript = true
	}

	if !haveRegion && !haveScript {
		return lang
	}

	var modifier string
	if haveRegion && haveScript {
		modifier = strings.ReplaceAll(l.DisplayPattern.Separator, displayToken0, script.String())
		modifier = strings.ReplaceAll(modifier, displayToken1, region.String())
	} else {
		if haveRegion {
			//special case - see if there's an l.Languages entry for base_Region
			baseReg := strings.ReplaceAll(tag.String(), "-", "_")
			if exact, ok := l.Languages[baseReg]; ok {
				return exact
			}

			//didn't have the special case. get the region name from territories
			if regionName, ok := l.Territories[region.String()]; ok {
				modifier = regionName
			}
		}
		if haveScript {
			modifier = script.String()
		}
	}

	pattern := strings.ReplaceAll(l.DisplayPattern.Pattern, displayToken0, lang)
	pattern = strings.ReplaceAll(pattern, displayToken1, modifier)

	return pattern
}
