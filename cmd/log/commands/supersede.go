package commands

import (
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-meta/cmd/log/forms"
	"github.com/launchdarkly/sdk-meta/lib/logs"
)

func RunSupersedeCommand() {
	var supersedeParams forms.SupersedeFormData
	err := logs.UpdateCodes(func(codes *logs.LdLogCodesJson) error {
		form := forms.NewSupersedeForm(codes, &supersedeParams)
		err := form.Run()
		if err != nil {
			return err
		}

		err = logs.SupersedeCode(codes, supersedeParams.SupersededCode, supersedeParams.NewCode, supersedeParams.Reason)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		return
	}
	fmt.Println("Superseded code")
}
