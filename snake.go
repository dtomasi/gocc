package gocc

import (
	"regexp"
	"strings"
)

func toSnakeCase(s string) string {
	switch detectCaseStyle(s) {
	case StyleKebabCase, StyleUpperKebabCase:
		return strings.ToLower(strings.ReplaceAll(s, "-", "_"))
	case StyleDotNotation, StyleUpperDotNotation:
		return strings.ToLower(strings.ReplaceAll(s, ".", "_"))
	case StyleCamelCase, StylePascalCase:
		return pascalOrCamelToSnakeCase(s)
	case StyleUpperSnakeCase:
		return strings.ToLower(s)
	}
	return s
}

func pascalOrCamelToSnakeCase(s string) string {
	var (
		matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	)
	s = matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	s = matchAllCap.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(s)
}
