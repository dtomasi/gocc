package gocc

import "fmt"

const (
	DelimiterSnakeCase   = "_"
	DelimiterKebabCase   = "-"
	DelimiterDotNotation = "."
)

// Check Regular expressions
var (
	regexCheckSnackCase        = buildCustomDelimiterRegex(DelimiterSnakeCase)
	regexCheckUpperSnakeCase   = buildUpperCustomDelimiterRegex(DelimiterSnakeCase)
	regexCheckPascalCase       = `^[A-Z][a-z0-9]+([A-Za-z0-9]+)*$`
	regexCheckCamelCase        = `^[a-z0-9]+([A-Za-z0-9]+)*$`
	regexCheckKebabCase        = buildCustomDelimiterRegex(DelimiterKebabCase)
	regexCheckUpperKebabCase   = buildUpperCustomDelimiterRegex(DelimiterKebabCase)
	regexCheckDotNotation      = buildCustomDelimiterRegex(DelimiterDotNotation)
	regexCheckUpperDotNotation = buildUpperCustomDelimiterRegex(DelimiterDotNotation)
)

var regexReservedChars = []string{`.`, `|`, `*`, `+`, `?`, `^`, `$`, `/`, `-`, `\`}

func buildCustomDelimiterRegex(delimiter string) string {
	if delimiterIsReservedChar(delimiter) {
		delimiter = `\` + delimiter
	}

	return fmt.Sprintf(`^[a-z][a-z0-9]+(%s[a-z0-9]+)*$`, delimiter)
}

func buildUpperCustomDelimiterRegex(delimiter string) string {
	if delimiterIsReservedChar(delimiter) {
		delimiter = `\` + delimiter
	}

	return fmt.Sprintf(`^[A-Z][A-Z0-9]+(%s[A-Z0-9]+)*$`, delimiter)
}

func delimiterIsReservedChar(delimiter string) bool {
	for _, a := range regexReservedChars {
		if a == delimiter {
			return true
		}
	}
	return false
}
