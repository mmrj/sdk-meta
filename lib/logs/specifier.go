package logs

import "regexp"

const specifierFormat = "^[a-zA-Z]+$"

func ValidSpecifierName(s string) bool {
	matched, err := regexp.MatchString(specifierFormat, s)
	if err != nil {
		return false
	}
	return matched
}

type SpecifierType string

var (
	SpecifierTypeClass     SpecifierType = "class"
	SpecifierTypeSystem    SpecifierType = "system"
	SpecifierTypeCondition SpecifierType = "condition"
)
