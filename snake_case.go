package gocc

import "strings"

func toSnakeCaseFromCaseStyle(from CaseStyle, s string) string {

	switch from {

	case StylePascalCase, StyleCamelCase:
		return pascalOrCamelCaseToSnakeCase(s)
	case StyleKebabCase, StyleUpperKebabCase:
		return kebabCaseToSnakeCase(s)
	case StyleDotNotation, StyleUpperDotNotation:
		return dotNotationToSnakeCase(s)
	case StyleUpperSnakeCase:
		return strings.ToLower(s)
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
