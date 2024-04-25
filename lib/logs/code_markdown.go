package logs

import (
	"errors"
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-meta/lib/collections"
	"github.com/launchdarkly/sdk-meta/lib/markdown"
)

type codeMarkdownWriter struct {
	writer *markdown.Writer
	codes  *LdLogCodesJson
}

func GenerateMarkdown(codes *LdLogCodesJson, outPath string) error {
	cw := codeMarkdownWriter{
		writer: markdown.NewWriter(1),
		codes:  codes,
	}

	cw.writeIntroduction()

	cw.writer.WriteSection("Codes", func() {
		collections.MapForEachOrdered(codes.Systems, func(systemName string, system System) {
			cw.writeSystem(systemName)
		})
	})

	return cw.writer.Save(outPath)
}

func GenMarkdownCondition(codes *LdLogCodesJson, code string) (string, error) {
	cw := codeMarkdownWriter{
		writer: markdown.NewWriter(1),
		codes:  codes,
	}

	condition, present := cw.codes.Conditions[code]
	if !present {
		return "", errors.New("code does not exist")
	}
	cw.writeCondition(code, condition)
	return cw.writer.String(), nil
}

func (cw *codeMarkdownWriter) writeIntroduction() {
	introductionPath := "logs/doc/fragments/introduction.md"
	if _, err := os.Stat(introductionPath); err == nil {
		fmt.Println("Fragment exists for:", introductionPath)
		err = cw.writer.AppendMarkdown(introductionPath)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("failed to write fragment for introduction, error: %w", err))
		}
	} else {
		fmt.Println("No fragment exists for:", introductionPath)
	}
}

func (cw *codeMarkdownWriter) writeCondition(code string, condition Condition) {
	sectionName := fmt.Sprintf("%s - %s", code, condition.Name)
	cw.writer.WriteSection(sectionName, func() {
		sysName, _, _ := collections.MapFind(cw.codes.Systems, func(s string, system System) bool {
			return system.Specifier == condition.System
		})
		className, _, _ := collections.MapFind(cw.codes.Classes, func(s string, class Class) bool {
			return class.Specifier == condition.Class
		})
		fmt.Println("Writing condition:", condition.Name, "of system", sysName, "with class", className)

		cw.writer.WriteLn(condition.Description)
		cw.writer.WriteBlankLn()

		cw.writer.WriteTableHeader("code", "system", "class")
		cw.writer.WriteTableRow(code, sysName, className)

		cw.writer.WriteBlankLn()
		cw.writer.WriteSection("Message", func() {
			cw.writer.WriteLn(fmt.Sprintf("`%s`", condition.Message.Parameterized))
			cw.writer.WriteBlankLn()

			if len(condition.Message.Parameters) != 0 {
				cw.writer.WriteTableHeader("parameter", "description")
				for paramName, paramDesc := range condition.Message.Parameters {
					cw.writer.WriteTableRow(paramName, paramDesc)
				}
			}

		})

		fragment := fmt.Sprintf("logs/doc/fragments/%s_%s_%s.md", sysName, className, condition.Name)

		if _, err := os.Stat(fragment); err == nil {
			fmt.Println("Fragment exists for:", fragment)
			err = cw.writer.AppendMarkdown(fragment)
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("failed to write fragment for code: %s error: %w", condition.Name, err))
			}
		} else {
			fmt.Println("No fragment exists for:", fragment)
		}
	})
}

func (cw *codeMarkdownWriter) writeSystem(name string) {
	cw.writer.WriteSection(name, func() {
		system := cw.codes.Systems[name]
		cw.writer.WriteLn(system.Description)
		cw.writer.WriteBlankLn()

		conditions := collections.MapFilter(cw.codes.Conditions, func(condition Condition) bool {
			return condition.System == system.Specifier
		})

		collections.MapForEachOrdered(conditions, func(code string, condition Condition) {
			cw.writeCondition(code, condition)
		})
	})
}
