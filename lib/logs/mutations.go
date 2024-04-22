package logs

import (
	"fmt"
)

func AddSystem(codes *LdLogCodesJson, name string, description string) error {
	err := ValidateSystemName(name, codes)
	if err != nil {
		return err
	}

	maxSpecifier := -1.0
	for _, system := range codes.Systems {
		if system.Specifier > maxSpecifier {
			maxSpecifier = system.Specifier
		}
	}
	newSpecifier := maxSpecifier + 1
	codes.Systems[name] = System{
		Description: description,
		Specifier:   newSpecifier,
	}
	return nil
}

func AddClass(codes *LdLogCodesJson, name string, description string) error {
	err := ValidateClassName(name, codes)
	if err != nil {
		return err
	}

	maxSpecifier := -1.0
	for _, system := range codes.Classes {
		if system.Specifier > maxSpecifier {
			maxSpecifier = system.Specifier
		}
	}
	newSpecifier := maxSpecifier + 1
	codes.Classes[name] = Class{
		Description: description,
		Specifier:   newSpecifier,
	}
	return nil
}

func AddCode(codes *LdLogCodesJson, className string, systemName string, conditionName string, description string, message Message) error {
	system, systemPresent := codes.Systems[systemName]

	if !systemPresent {
		return fmt.Errorf("the system class does not exist. Please choose an existing system or create a new system")
	}

	class, classPresent := codes.Classes[className]

	if !classPresent {
		return fmt.Errorf("the specified class does not exist. Please choose an existing class or create a new class")
	}

	_, present := codes.Conditions[conditionName]
	if present {
		return fmt.Errorf("condition name already exists. Please choose a new name or using the existing specifier")
	}

	maxSpecifier := -1.0
	for _, condition := range codes.Conditions {
		if condition.Specifier > maxSpecifier {
			maxSpecifier = condition.Specifier
		}
	}
	newSpecifier := maxSpecifier + 1
	codes.Conditions[conditionName] = Condition{
		Description: description,
		Specifier:   newSpecifier,
		Class:       class.Specifier,
		System:      system.Specifier,
		Message:     message,
	}

	return nil
}
