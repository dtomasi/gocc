package gocc

import "strings"

func toKebabCaseFromCaseStyle(from CaseStyle, s string) string {
	switch from {
	case StyleUnknown, StyleKebabCase:
		panic("cannot convert from StyleUnknown or same as input style")
	case StylePascalCase, StyleCamelCase:
		return pascalOrCamelCaseToKebabCase(s)
	case StyleSnakeCase:
		return snakeCaseToKebabCase(s)
	case StyleDotNotation:
		return dotNotationToKebabCase(s)
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
