package forms

import (
	"errors"

	"github.com/charmbracelet/huh"
	"github.com/launchdarkly/sdk-meta/lib/logs"
)

type DeprecateFormData struct {
	Code   string
	Reason string
}

func NewDeprecateForm(codes *logs.LdLogCodesJson, deprecate *DeprecateFormData) *huh.Form {
	// Starting group decides which type of specifier to create.
	return huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Please enter the code to deprecate?").
			Value(&deprecate.Code).
			Validate(func(s string) error {
				if !logs.ValidateCode(s) {
					return errors.New("the code was not in the correct format")
				}
				for _, condition := range codes.Conditions {
					if s == logs.GetCode(condition) {
						return nil
					}
				}
				return errors.New("could not find an existing entry matching the code")
			}),
		huh.NewInput().
			Title("What is the reason the code is being deprecated?").
			Value(&deprecate.Reason),
	))
}
