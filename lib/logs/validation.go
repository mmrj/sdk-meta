package logs

import (
	"fmt"
	"regexp"

	"github.com/launchdarkly/sdk-meta/lib/collections"
)

const codeFormat = "^[0-9]+:[0-9]+:[0-9]+$"

func validateName(name string, specifierType SpecifierType, present bool) error {
	if present {
		return fmt.Errorf("%s name already exists. Please choose a new name or use the existing specifier", specifierType)
	}
	if !ValidSpecifierName(name) {
		return fmt.Errorf("the %s name must be composed of only upper and lowercase ASCII letters and may not be empty [a-zA-Z]+", specifierType)
	}
	return nil
}

func ValidateSystemName(name string, codes *LdLogCodesJson) error {
	_, present := codes.Systems[name]
	return validateName(name, SpecifierTypeSystem, present)
}

func ValidateClassName(name string, codes *LdLogCodesJson) error {
	_, present := codes.Classes[name]
	return validateName(name, SpecifierTypeSystem, present)
}

func ValidateConditionName(name string, system float64, class float64, codes *LdLogCodesJson) error {
	conditions := collections.MapFilter(codes.Conditions, func(condition Condition) bool {
		return condition.System == system && condition.Class == class
	})
	_, _, present := collections.MapFind(conditions, func(s string, condition Condition) bool {
		return condition.Name == name
	})
	return validateName(name, SpecifierTypeCondition, present)
}

func ValidateParameterizedMessageString(message string) error {
	_, err := ParseMessage(message)
	return err
}

func ValidateCode(code string) bool {
	matched, err := regexp.MatchString(codeFormat, code)
	if err != nil {
		return false
	}
	return matched
}
