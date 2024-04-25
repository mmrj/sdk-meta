package commands

import (
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-meta/cmd/log/forms"
	"github.com/launchdarkly/sdk-meta/lib/logs"
)

func RunNewCommand() {
	var specifierType logs.SpecifierType
	form := forms.NewSpecifierForm(&specifierType)
	err := form.Run()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error specifier form", err.Error())
		return
	}
	switch specifierType {
	case logs.SpecifierTypeSystem:
		runNewSystemCommand()
	case logs.SpecifierTypeClass:
		runNewClassCommand()
	case logs.SpecifierTypeCondition:
		runNewConditionCommand()
	}
}

func runNewClassCommand() {
	err := logs.UpdateCodes(func(codes *logs.LdLogCodesJson) error {
		var params forms.ClassFormData
		form := forms.NewClassForm(codes, &params)
		err := form.Run()
		if err != nil {
			return fmt.Errorf("error running new system form: %w", err)
		}
		err = logs.AddClass(codes, params.Name, params.Description)
		return err
	})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func runNewSystemCommand() {
	err := logs.UpdateCodes(func(codes *logs.LdLogCodesJson) error {
		var params forms.SystemFormData
		form := forms.NewSystemForm(codes, &params)
		err := form.Run()
		if err != nil {
			return fmt.Errorf("error running new system form: %w", err)
		}
		err = logs.AddSystem(codes, params.Name, params.Description)
		return err
	})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func runNewConditionCommand() {
	err := logs.UpdateCodes(func(codes *logs.LdLogCodesJson) error {
		var params forms.ConditionFormData
		form := forms.NewConditionFormPart1(codes, &params)
		err := form.Run()

		if err != nil {
			return fmt.Errorf("error running new condition form part 1: %w", err)
		}

		form = forms.NewConditionFormPart2(codes, &params)
		err = form.Run()
		if err != nil {
			return fmt.Errorf("error running new condition form part 2: %w", err)
		}

		parameters, err := logs.ParseMessage(params.MessageString)
		if err != nil {
			return fmt.Errorf("bad message string: %w", err)
		}

		message := logs.Message{
			Parameters:    map[string]string{},
			Parameterized: params.MessageString,
		}

		if len(parameters.Parameters) != 0 {
			var paramParams forms.ParameterFormData
			parametersForm := forms.NewMessageForm(codes, parameters.Parameters, &paramParams)
			err = parametersForm.Run()
			if err != nil {
				return fmt.Errorf("error running parameters form: %w", err)
			}
			for key, value := range paramParams.Descriptions {
				message.Parameters[key] = *value
			}
		}

		err = logs.AddCondition(codes, params.Class, params.System, params.Name, params.Description, message)
		if err == nil {
			fmt.Printf("The \"%s\" condition has been added.", params.Name)
		}
		return err
	})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}
