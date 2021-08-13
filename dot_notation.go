package gocc

import "strings"

func toDotNotationFromCaseStyle(from CaseStyle, s string) string {

	switch from {

	case StylePascalCase, StyleCamelCase:
		return pascalOrCamelCaseToDotNotation(s)
	case StyleKebabCase, StyleUpperKebabCase:
		return kebabCaseToDotNotation(s)
	case StyleSnakeCase, StyleUpperSnakeCase:
		return snakeCaseToDotNotation(s)
	case StyleUpperDotNotation:
		return strings.ToUpper(s)
	}

	return s
}

func pascalOrCamelCaseToDotNotation(s string) string {
	return pascalOrCamelToCustomDelimiter(s, DelimiterDotNotation)
}

func kebabCaseToDotNotation(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, DelimiterKebabCase, DelimiterDotNotation))
}

func snakeCaseToDotNotation(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, DelimiterSnakeCase, DelimiterDotNotation))
}