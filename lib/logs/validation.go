package logs

import (
	"fmt"
)

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

func ValidateConditionName(name string, codes *LdLogCodesJson) error {
	_, present := codes.Conditions[name]
	return validateName(name, SpecifierTypeCondition, present)
}

func ValidateParameterizedMessageString(message string) error {
	return nil
}
