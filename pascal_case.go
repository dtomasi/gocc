package gocc

import (
	"fmt"
	"regexp"
	"strings"
)

func toPascalCaseFromCaseStyle(from CaseStyle, s string) string {

	switch from {
	case StyleCamelCase:
		return ucFirst(s)
	case StyleSnakeCase, StyleUpperSnakeCase:
		return customDelimiterToPascalCase(s, DelimiterSnakeCase)
	case StyleKebabCase, StyleUpperKebabCase:
		return customDelimiterToPascalCase(s, DelimiterKebabCase)
	case StyleDotNotation, StyleUpperDotNotation:
		return customDelimiterToPascalCase(s, DelimiterDotNotation)
	}

	return s
}

func customDelimiterToPascalCase(s string, delimiter string) string {
	regexDelimiter := delimiter
	if delimiterIsReservedChar(delimiter) {
		regexDelimiter = `\` + delimiter
	}
	regexString := fmt.Sprintf(`(^[A-Za-z])|%s([A-Za-z])`, regexDelimiter)
	var r = regexp.MustCompile(regexString)
	return r.ReplaceAllStringFunc(strings.ToLower(s), func(s string) string {
		return strings.ToUpper(strings.Replace(s, delimiter, "", -1))
	})
}
