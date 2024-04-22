package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/launchdarkly/sdk-meta/lib/logs"
)

func NewSpecifierForm() (*huh.Form, *logs.SpecifierType) {
	var specifierType logs.SpecifierType
	// Starting group decides which type of specifier to create.
	return huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().
			Title("Choose the type of specifier to create").
			Options(
				huh.NewOption("system", string(logs.SpecifierTypeSystem)),
				huh.NewOption("class", string(logs.SpecifierTypeClass)),
				huh.NewOption("condition", string(logs.SpecifierTypeCondition)),
			).Value((*string)(&specifierType)),
	)), &specifierType
}
