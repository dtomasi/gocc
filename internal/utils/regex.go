package utils

import "fmt"

const (
	// well known delimiters.
	DelimiterSnakeCase   = "_"
	DelimiterKebabCase   = "-"
	DelimiterDotNotation = "."

	// regex to extract delimiter from given string.
	RegexDelimiterDetection = `^[a-z][a-z0-9].+([_\-\.\|\*\+\?\^\$\/\\,;])[a-z0-9]*$`
)

// Check Regular expressions.
var (
	RegexCheckSnackCase   = BuildCustomDelimiterRegex(DelimiterSnakeCase)
	RegexCheckPascalCase  = `^[A-Z][a-z0-9]+([A-Za-z0-9]+)*$`
	RegexCheckCamelCase   = `^[a-z0-9]+([A-Za-z0-9]+)*$`
	RegexCheckKebabCase   = BuildCustomDelimiterRegex(DelimiterKebabCase)
	RegexCheckDotNotation = BuildCustomDelimiterRegex(DelimiterDotNotation)
)

// RegexReservedChars reserved chars for regex patterns.
var RegexReservedChars = []string{`.`, `|`, `*`, `+`, `?`, `^`, `$`, `/`, `-`, `\`}

// BuildCustomDelimiterRegex generates a regular expression to check string contains given delimiter.
func BuildCustomDelimiterRegex(delimiter string) string {
	if DelimiterIsReservedChar(delimiter) {
		delimiter = `\` + delimiter
	}

	return fmt.Sprintf(`^[a-z][a-z0-9%s]*$`, delimiter)
}

// BuildUpperCustomDelimiterRegex same as buildCustomDelimiterRegex but for upper case strings.
func BuildUpperCustomDelimiterRegex(delimiter string) string {
	if DelimiterIsReservedChar(delimiter) {
		delimiter = `\` + delimiter
	}

	return fmt.Sprintf(`^[A-Z][A-Z0-9%s]*$`, delimiter)
}

// DelimiterIsReservedChar checks if delimiter is in reserved chars slice.
func DelimiterIsReservedChar(delimiter string) bool {
	for _, a := range RegexReservedChars {
		if a == delimiter {
			return true
		}
	}

	return false
}
