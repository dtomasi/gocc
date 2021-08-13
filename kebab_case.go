package gocc

import "strings"

func toKebabCaseFromCaseStyle(from CaseStyle, s string) string {

	switch from {

	case StylePascalCase, StyleCamelCase:
		return pascalOrCamelCaseToKebabCase(s)
	case StyleSnakeCase, StyleUpperSnakeCase:
		return snakeCaseToKebabCase(s)
	case StyleDotNotation, StyleUpperDotNotation:
		return dotNotationToKebabCase(s)
	case StyleUpperKebabCase:
		return strings.ToUpper(s)
	}

	return s
}

func pascalOrCamelCaseToKebabCase(s string) string {
	return pascalOrCamelToCustomDelimiter(s, DelimiterKebabCase)
}

func snakeCaseToKebabCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, DelimiterSnakeCase, DelimiterKebabCase))
}

func dotNotationToKebabCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, DelimiterDotNotation, DelimiterKebabCase))
}