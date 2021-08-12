package gocc

import (
	"regexp"
	"strings"
)

func toCamelCase(s string) string {
	switch detectCaseStyle(s) {
	case StyleUpperSnakeCase, StyleUpperKebabCase:
		return toCamelCase(strings.ToLower(s))
	case StyleSnakeCase:
		var r = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")
		return lcFirst(r.ReplaceAllStringFunc(s, func(s string) string {
			return strings.ToUpper(strings.Replace(s, "_", "", -1))
		}))
	case StylePascalCase:
		return lcFirst(s)
	case StyleKebabCase:
		var r = regexp.MustCompile("(^[A-Za-z])|-([A-Za-z])")
		return lcFirst(r.ReplaceAllStringFunc(s, func(s string) string {
			return strings.ToUpper(strings.Replace(s, "-", "", -1))
		}))
	}
	return s
}
