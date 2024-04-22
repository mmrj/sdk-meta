package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/launchdarkly/sdk-meta/lib/logs"
)

type ConditionFormData struct {
	Description   string
	Name          string
	MessageString string
	System        string
	Class         string
}

func NewConditionForm(codes *logs.LdLogCodesJson, condition *ConditionFormData) *huh.Form {
	var systemOptions []huh.Option[string]
	var classOptions []huh.Option[string]

	for systemName, _ := range codes.Systems {
		systemOptions = append(systemOptions, huh.NewOption(systemName, systemName))
	}

	for className, _ := range codes.Classes {
		classOptions = append(classOptions, huh.NewOption(className, className))
	}

	// Starting group decides which type of specifier to create.
	return huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("The name of the condition?").
			Value(&condition.Name).
			Validate(func(s string) error {
				return logs.ValidateConditionName(s, codes)
			}),
		huh.NewInput().
			Title("Please describe the condition?").
			Value(&condition.Description),
		huh.NewSelect[string]().Title("Select system:").Options(systemOptions...).Value(&condition.System),
		huh.NewSelect[string]().Title("Select class:").Options(classOptions...).Value(&condition.Class),
		huh.NewInput().
			Title("Parameterized message string for the condition:").
			Value(&condition.MessageString).
			Validate(func(s string) error {
				return logs.ValidateParameterizedMessageString(condition.MessageString)
			}),
	))
}
