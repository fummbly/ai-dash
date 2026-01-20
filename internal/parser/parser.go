// Package parser for creating utility functions for reusable code
package parser

import (
	"strings"
)

func ConvertBold(original string) string {
	count := strings.Count(original, "**")

	if count%2 != 0 {
		count -= 1
	}

	originalTemp := original
	newString := ""

	for i := range count {
		index := strings.Index(originalTemp, "**")
		newString += originalTemp[:index]

		if i%2 == 0 {
			newString += "<strong>"
		} else {
			newString += "</strong>"
		}

		originalTemp = originalTemp[index+2:]
	}

	return newString + originalTemp
}

func CovertNewLine(original string) string {
	return strings.ReplaceAll(original, "\n", "<br>")
}

func ConvertLineBreak(original string) string {
	return strings.ReplaceAll(original, "---", "<hr>")
}

func ConvertMarkdown(original string) string {
	converted := ConvertBold(original)
	converted = ConvertLineBreak(converted)
	converted = CovertNewLine(converted)

	return converted
}
