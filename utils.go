package gocc

import (
	"fmt"
	"regexp"
	"strings"
)

// ucFirst is a shortcut for strings.Title
// capitalizes the first char
func ucFirst(s string) string {
	return strings.Title(s)
}

// lcFirst lowers the first char
func lcFirst(s string) string {
	r := regexp.MustCompile(`^(\b[A-Z])`)
	return r.ReplaceAllStringFunc(s, func(s string) string {
		return strings.ToLower(s)
	})
}

func toCustomDelimiter(s string, delimiter string) string {

	switch detectCaseStyle(s) {
	case StyleUpperSnakeCase, StyleUpperKebabCase, StyleUpperDotNotation:
		return toCustomDelimiter(strings.ToLower(s), delimiter)
	case StyleSnakeCase:
		return strings.ToLower(strings.ReplaceAll(s, DelimiterSnakeCase, delimiter))
	case StyleKebabCase:
		return strings.ToLower(strings.ReplaceAll(s, DelimiterKebabCase, delimiter))
	case StyleDotNotation:
		return strings.ToLower(strings.ReplaceAll(s, DelimiterDotNotation, delimiter))
	case StylePascalCase, StyleCamelCase:
		return pascalOrCamelToCustomDelimiter(s, delimiter)
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

func customDelimiterToPascalCase(s string, delimiter string) string {
	regexDelimiter := delimiter
	if delimiterIsReservedChar(delimiter) {
		regexDelimiter = `\` + delimiter
	}
	regexString := fmt.Sprintf(`(^[A-Za-z])|%s([A-Za-z])`, regexDelimiter)
	var r = regexp.MustCompile(regexString)
	return r.ReplaceAllStringFunc(s, func(s string) string {
		return strings.ToUpper(strings.Replace(s, delimiter, "", -1))
	})
}

func customDelimiterToCamelCase(s string, delimiter string) string {
	return lcFirst(customDelimiterToPascalCase(s, delimiter))
}

func toPascalCase(s string) string {
	switch detectCaseStyle(s) {
	case StyleUpperSnakeCase, StyleUpperKebabCase, StyleUpperDotNotation:
		return toPascalCase(strings.ToLower(s))
	case StyleCamelCase:
		return ucFirst(s)
	case StyleSnakeCase:
		return customDelimiterToPascalCase(s, DelimiterSnakeCase)
	case StyleKebabCase:
		return customDelimiterToPascalCase(s, DelimiterKebabCase)
	case StyleDotNotation:
		return customDelimiterToPascalCase(s, DelimiterDotNotation)
	}
	return s
}

func toCamelCase(s string) string {
	switch detectCaseStyle(s) {
	case StyleUpperSnakeCase, StyleUpperKebabCase, StyleUpperDotNotation:
		return toCamelCase(strings.ToLower(s))
	case StyleSnakeCase:
		return customDelimiterToCamelCase(s, DelimiterSnakeCase)
	case StylePascalCase:
		return lcFirst(s)
	case StyleKebabCase:
		return customDelimiterToCamelCase(s, DelimiterKebabCase)
	case StyleDotNotation:
		return customDelimiterToCamelCase(s, DelimiterDotNotation)
	}
	return s
}
