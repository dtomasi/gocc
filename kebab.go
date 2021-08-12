package gocc

import (
	"regexp"
	"strings"
)

func toKebabCase(s string) string {
	switch detectCaseStyle(s) {
	case StyleSnakeCase, StyleUpperSnakeCase:
		return strings.ToLower(strings.ReplaceAll(s, "_", "-"))
	case StylePascalCase, StyleCamelCase:
		return pascalOrCamelToKebabCase(s)
	case StyleUpperKebabCase:
		return strings.ToLower(s)
	}

	return s
}

func pascalOrCamelToKebabCase(s string) string {
	var (
		matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	)
	s = matchFirstCap.ReplaceAllString(s, "${1}-${2}")
	s = matchAllCap.ReplaceAllString(s, "${1}-${2}")
	return strings.ToLower(s)
}
