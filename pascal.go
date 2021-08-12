package gocc

import (
	"regexp"
	"strings"
)

func toPascalCase(s string) string {
	switch detectCaseStyle(s) {
	case StyleUpperSnakeCase, StyleUpperKebabCase, StyleUpperDotNotation:
		return toPascalCase(strings.ToLower(s))
	case StyleSnakeCase:
		var r = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")
		return r.ReplaceAllStringFunc(s, func(s string) string {
			return strings.ToUpper(strings.Replace(s, "_", "", -1))
		})
	case StyleCamelCase:
		return ucFirst(s)
	case StyleKebabCase:
		var r = regexp.MustCompile("(^[A-Za-z])|-([A-Za-z])")
		return r.ReplaceAllStringFunc(s, func(s string) string {
			return strings.ToUpper(strings.Replace(s, "-", "", -1))
		})
	case StyleDotNotation:
		var r = regexp.MustCompile("(^[A-Za-z])|.([A-Za-z])")
		return r.ReplaceAllStringFunc(s, func(s string) string {
			return strings.ToUpper(strings.Replace(s, ".", "", -1))
		})
	}
	return s
}
