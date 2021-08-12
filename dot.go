package gocc

import (
	"regexp"
	"strings"
)

func toDotNotation(s string) string {
	switch detectCaseStyle(s) {
	case StyleKebabCase, StyleUpperKebabCase:
		return strings.ToLower(strings.ReplaceAll(s, "-", "."))
	case StyleSnakeCase, StyleUpperSnakeCase:
		return strings.ToLower(strings.ReplaceAll(s, "_", "."))
	case StyleCamelCase, StylePascalCase:
		return pascalOrCamelToDotNotation(s)
	case StyleUpperDotNotation:
		return strings.ToLower(s)
	}
	return s
}

func pascalOrCamelToDotNotation(s string) string {
	var (
		matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	)
	s = matchFirstCap.ReplaceAllString(s, "${1}.${2}")
	s = matchAllCap.ReplaceAllString(s, "${1}.${2}")
	return strings.ToLower(s)
}
