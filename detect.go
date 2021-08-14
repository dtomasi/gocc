package gocc

import (
	"regexp"
	"strings"
)

// detectCaseStyle returns the detected style as int (CaseStyle)
// see style.go
func detectCaseStyle(s string) CaseStyle {

	switch true {
	case IsPascalCase(s):
		return StylePascalCase
	case IsCamelCase(s):
		return StyleCamelCase
	}

	delimiter, found := detectDelimitedString(s)
	if !found {
		return StyleUnknown
	}

	switch delimiter {
	case DelimiterSnakeCase:
		return StyleSnakeCase
	case DelimiterKebabCase:
		return StyleKebabCase
	case DelimiterDotNotation:
		return StyleDotNotation
	}

	return StyleUnknown
}

func detectDelimitedString(s string) (string, bool) {
	r := regexp.MustCompile(regexDelimiterDetection)
	subMatch := r.FindStringSubmatch(s)
	if len(subMatch) > 0 {
		// this is a required validation step to check if a string only contains detected delimiter
		if IsCustomDelimiter(s, subMatch[1]) {
			return subMatch[1], true
		}
	}
	return "", false
}

// IsSnakeCase detects if a string is snake case style
func IsSnakeCase(s string) bool {
	r := regexp.MustCompile(regexCheckSnackCase)
	return r.MatchString(strings.TrimSpace(s))
}

// IsPascalCase detects if a string is pascal case style
func IsPascalCase(s string) bool {
	r := regexp.MustCompile(regexCheckPascalCase)
	return r.MatchString(strings.TrimSpace(s))
}

// IsCamelCase detects if a string is camel case style
func IsCamelCase(s string) bool {
	r := regexp.MustCompile(regexCheckCamelCase)
	return r.MatchString(strings.TrimSpace(s))
}

// IsKebabCase detects if a string is kebab case style
func IsKebabCase(s string) bool {
	r := regexp.MustCompile(regexCheckKebabCase)
	return r.MatchString(strings.TrimSpace(s))
}

// IsDotNotation detects if a string is dot notation style
func IsDotNotation(s string) bool {
	r := regexp.MustCompile(regexCheckDotNotation)
	return r.MatchString(strings.TrimSpace(s))
}

// IsCustomDelimiter detects if a string is separated by a given custom delimiter
func IsCustomDelimiter(s string, delimiter string) bool {
	r := regexp.MustCompile(buildCustomDelimiterRegex(delimiter))
	return r.MatchString(strings.TrimSpace(s))
}

// IsUpperCustomDelimiter detects if a string is separated by a given custom delimiter as upper case string
func IsUpperCustomDelimiter(s string, delimiter string) bool {
	r := regexp.MustCompile(buildUpperCustomDelimiterRegex(delimiter))
	return r.MatchString(strings.TrimSpace(s))
}
