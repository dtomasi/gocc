package gocc

import "strings"

func toSnakeCaseFromCaseStyle(from CaseStyle, s string) string {
	switch from {
	case StyleUnknown:
		panic("cannot convert from StyleUnknown")
	case StyleSnakeCase:
		return s
	case StylePascalCase, StyleCamelCase:
		return pascalOrCamelCaseToSnakeCase(s)
	case StyleKebabCase:
		return kebabCaseToSnakeCase(s)
	case StyleDotNotation:
		return dotNotationToSnakeCase(s)
	}

	return s
}

func pascalOrCamelCaseToSnakeCase(s string) string {
	return pascalOrCamelToCustomDelimiter(s, DelimiterSnakeCase)
}

func kebabCaseToSnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, DelimiterKebabCase, DelimiterSnakeCase))
}

func dotNotationToSnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, DelimiterDotNotation, DelimiterSnakeCase))
}
