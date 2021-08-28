package gocc

import (
	"regexp"
	"strings"
)

func toCustomDelimiterFromCaseStyle(from CaseStyle, s string, delimiter string) string {
	switch from {
	case StyleUnknown:
		panic("cannot convert from StyleUnknown")
	case StylePascalCase, StyleCamelCase:
		return pascalOrCamelToCustomDelimiter(s, delimiter)
	case StyleSnakeCase:
		return strings.ToLower(strings.ReplaceAll(s, DelimiterSnakeCase, delimiter))
	case StyleKebabCase:
		return strings.ToLower(strings.ReplaceAll(s, DelimiterKebabCase, delimiter))
	case StyleDotNotation:
		return strings.ToLower(strings.ReplaceAll(s, DelimiterDotNotation, delimiter))
	}

	return s
}

func pascalOrCamelToCustomDelimiter(s string, delimiter string) string {
	s = lcFirst(s)
	r := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	s = r.ReplaceAllStringFunc(s, func(s string) string {
		return strings.Join(strings.Split(s, ""), delimiter)
	})

	return strings.ToLower(s)
}
