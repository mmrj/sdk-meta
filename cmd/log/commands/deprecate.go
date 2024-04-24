package commands

import (
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-meta/cmd/log/forms"
	"github.com/launchdarkly/sdk-meta/lib/logs"
)

func RunDeprecateCommand() {
	var deprecateParams forms.DeprecateFormData
	err := logs.UpdateCodes(func(codes *logs.LdLogCodesJson) error {
		form := forms.NewDeprecateForm(codes, &deprecateParams)
		err := form.Run()
		if err != nil {
			return err
		}

		err = logs.DeprecateCode(codes, deprecateParams.Code, deprecateParams.Reason)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		return
	}
	fmt.Println("Deprecated code")
}
