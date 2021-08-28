package gocc

import "fmt"

const (
	// well known delimiters.
	DelimiterSnakeCase   = "_"
	DelimiterKebabCase   = "-"
	DelimiterDotNotation = "."

	// regex to extract delimiter from given string.
	regexDelimiterDetection = `^[a-z][a-z0-9].+([_\-\.\|\*\+\?\^\$\/\\,;])[a-z0-9]*$`
)

// Check Regular expressions.
var (
	regexCheckSnackCase   = buildCustomDelimiterRegex(DelimiterSnakeCase)
	regexCheckPascalCase  = `^[A-Z][a-z0-9]+([A-Za-z0-9]+)*$`
	regexCheckCamelCase   = `^[a-z0-9]+([A-Za-z0-9]+)*$`
	regexCheckKebabCase   = buildCustomDelimiterRegex(DelimiterKebabCase)
	regexCheckDotNotation = buildCustomDelimiterRegex(DelimiterDotNotation)
)

// reserved chars for regex patterns.
var regexReservedChars = []string{`.`, `|`, `*`, `+`, `?`, `^`, `$`, `/`, `-`, `\`}

// buildCustomDelimiterRegex generates a regular expression to check string contains given delimiter.
func buildCustomDelimiterRegex(delimiter string) string {
	if delimiterIsReservedChar(delimiter) {
		delimiter = `\` + delimiter
	}

	return fmt.Sprintf(`^[a-z][a-z0-9%s]*$`, delimiter)
}

// buildUpperCustomDelimiterRegex same as buildCustomDelimiterRegex but for upper case strings.
func buildUpperCustomDelimiterRegex(delimiter string) string {
	if delimiterIsReservedChar(delimiter) {
		delimiter = `\` + delimiter
	}

	return fmt.Sprintf(`^[A-Z][A-Z0-9%s]*$`, delimiter)
}

// delimiterIsReservedChar checks if delimiter is in reserved chars slice.
func delimiterIsReservedChar(delimiter string) bool {
	for _, a := range regexReservedChars {
		if a == delimiter {
			return true
		}
	}

	return false
}
