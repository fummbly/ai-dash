package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoldParse(t *testing.T) {
	type BoldConversionTests struct {
		TestName     string
		InputText    string
		ExpectedText string
	}

	tests := []BoldConversionTests{
		{
			TestName:     "Basic bold conversion test",
			InputText:    "Testing bold **text** and this is **another** word",
			ExpectedText: "Testing bold <strong>text</strong> and this is <strong>another</strong> word",
		},
		{
			TestName:     "AI response test",
			InputText:    "* **Testing** bold **text** and this is **another** word",
			ExpectedText: "* <strong>Testing</strong> bold <strong>text</strong> and this is <strong>another</strong> word",
		},
		{
			TestName:     "Unclosed bold test",
			InputText:    "Testing bold **text** and this is **another word",
			ExpectedText: "Testing bold <strong>text</strong> and this is **another word",
		},
		{
			TestName:     "Single unclosed bold test",
			InputText:    "Testing bold **test and other stuff",
			ExpectedText: "Testing bold **test and other stuff",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			result := ConvertBold(test.InputText)
			assert.Equal(t, test.ExpectedText, result, "Got different string than expected")
		})
	}
}

func TestConvertMarkdown(t *testing.T) {
	type MarkdownConversionTests struct {
		TestName     string
		InputText    string
		ExpectedText string
	}

	tests := []MarkdownConversionTests{
		{
			TestName:     "Basic Markdown conversion test",
			InputText:    "This is a test of the **Markdown** conversion\nThis is for testing `ai` responses *of* different types\nAll is usefull for using things like this [https://test.com](Test thing)---this is it",
			ExpectedText: "This is a test of the <strong>Markdown</strong> conversion<br>This is for testing `ai` responses *of* different types<br>All is usefull for using things like this [https://test.com](Test thing)<hr>this is it",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			result := ConvertMarkdown(test.InputText)
			assert.Equal(t, test.ExpectedText, result, "Got different string than expected")
		})
	}
}
