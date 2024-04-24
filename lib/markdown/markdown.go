package markdown

import (
	"fmt"
	"os"
	"strings"

	"github.com/launchdarkly/sdk-meta/lib/collections"
	"github.com/launchdarkly/sdk-meta/lib/logs"
)

type markdownWriter struct {
	codes   *logs.LdLogCodesJson
	depth   int
	builder *strings.Builder
}

func (writer *markdownWriter) writeTableHeader(headers ...string) {
	writer.writeTableRow(headers...)
	writer.write("|-")
	for idx, header := range headers {
		writer.write(strings.Repeat("-", len(header)))

		if idx != len(headers)-1 {
			writer.write("-|-")
		}
	}
	writer.write("-|")
	writer.writeBlank()
}

func (writer *markdownWriter) writeTableRow(items ...string) {
	writer.write("| ")
	writer.write(strings.Join(items, " | "))
	writer.write(" |")
	writer.writeBlank()
}

func (writer *markdownWriter) writeCondition(name string, condition logs.Condition) {
	writer.writeSection(name, func() {
		sysName, _, _ := collections.MapFind(writer.codes.Systems, func(s string, system logs.System) bool {
			return system.Specifier == condition.System
		})
		clsName, _, _ := collections.MapFind(writer.codes.Classes, func(s string, class logs.Class) bool {
			return class.Specifier == condition.Class
		})

		writer.writeLn(condition.Description)
		writer.writeBlank()

		code := fmt.Sprintf("%d:%d:%d", int(condition.System), int(condition.Class), int(condition.Specifier))
		writer.writeTableHeader("code", "system", "class")
		writer.writeTableRow(code, sysName, clsName)

		writer.writeBlank()
		writer.writeSection("Message", func() {
			writer.writeLn(fmt.Sprintf("`%s`", condition.Message.Parameterized))
			writer.writeBlank()

			if len(condition.Message.Parameterized) != 0 {
				writer.writeTableHeader("parameter", "description")
				for paramName, paramDesc := range condition.Message.Parameters {
					writer.writeTableRow(paramName, paramDesc)
				}
			}

		})
	})
}

func (writer *markdownWriter) writeSystem(name string) {
	writer.writeSection(name, func() {
		system := writer.codes.Systems[name]
		writer.writeLn(system.Description)
		writer.writeBlank()

		conditions := collections.MapFilter(writer.codes.Conditions, func(condition logs.Condition) bool {
			return condition.System == system.Specifier
		})

		collections.MapForEachOrdered(conditions, func(condName string, condition logs.Condition) {
			writer.writeCondition(condName, condition)
		})
	})
}

func (writer *markdownWriter) writeSection(title string, content func()) {
	writer.depth += 1
	writer.write(strings.Repeat("#", writer.depth))
	writer.write(" ")
	writer.writeLn(title)
	writer.writeBlank()
	defer func() {
		writer.depth -= 1
	}()

	content()
	writer.writeBlank()
}

func (writer *markdownWriter) write(text string) {
	// Builder never fails a write. https://pkg.go.dev/strings#Builder.Write
	_, _ = writer.builder.WriteString(text)
}

func (writer *markdownWriter) writeLn(text string) {
	writer.write(text)
	writer.writeBlank()
}

func (writer *markdownWriter) writeBlank() {
	writer.write("\n")
}

func (writer *markdownWriter) save(outPath string) error {
	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	_, err = file.WriteString(writer.builder.String())
	return err
}

func GenerateMarkdown(codes *logs.LdLogCodesJson, outPath string) error {
	writer := markdownWriter{codes: codes, builder: &strings.Builder{}}

	collections.MapForEachOrdered(codes.Systems, func(systemName string, system logs.System) {
		writer.writeSystem(systemName)
	})

	return writer.save(outPath)
}
