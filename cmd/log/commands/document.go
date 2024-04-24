package commands

import (
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-meta/lib/logs"
	"github.com/launchdarkly/sdk-meta/lib/markdown"
)

func RunDocumentCommand() {
	err := logs.WithCodes(func(codes *logs.LdLogCodesJson) error {
		return markdown.GenerateMarkdown(codes, "logs/doc/codes.md")
	})

	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
	} else {
		println("Generated documentation in logs/doc/codes.md")
	}
}
