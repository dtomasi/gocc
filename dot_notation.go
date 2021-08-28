package gocc

import "strings"

func toDotNotationFromCaseStyle(from CaseStyle, s string) string {
	switch from {
	case StyleUnknown, StyleDotNotation:
		panic("cannot convert from StyleUnknown or same as input style")
	case StylePascalCase, StyleCamelCase:
		return pascalOrCamelCaseToDotNotation(s)
	case StyleKebabCase:
		return kebabCaseToDotNotation(s)
	case StyleSnakeCase:
		return snakeCaseToDotNotation(s)
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
