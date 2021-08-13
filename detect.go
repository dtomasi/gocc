package gocc

import (
	"regexp"
	"strings"
)

// detectCaseStyle returns the detected style as int (CaseStyle)
// see style.go
func detectCaseStyle(s string) CaseStyle {

	switch true {
	case IsSnakeCase(s):
		return StyleSnakeCase
	case IsUpperSnakeCase(s):
		return StyleUpperSnakeCase
	case IsPascalCase(s):
		return StylePascalCase
	case IsCamelCase(s):
		return StyleCamelCase
	case IsKebabCase(s):
		return StyleKebabCase
	case IsUpperKebabCase(s):
		return StyleUpperKebabCase
	case IsDotNotation(s):
		return StyleDotNotation
	case IsUpperDotNotation(s):
		return StyleUpperDotNotation
	}

	return StyleUnknown
}

// IsSnakeCase detects if a string is snake case style
func IsSnakeCase(s string) bool {
	r := regexp.MustCompile(regexCheckSnackCase)
	return r.MatchString(strings.TrimSpace(s))
}

// IsUpperSnakeCase detects if a string is snake case style
func IsUpperSnakeCase(s string) bool {
	r := regexp.MustCompile(regexCheckUpperSnakeCase)
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

// IsUpperKebabCase detects if a string is kebab case style
func IsUpperKebabCase(s string) bool {
	r := regexp.MustCompile(regexCheckUpperKebabCase)
	return r.MatchString(strings.TrimSpace(s))
}

// IsDotNotation detects if a string is dot notation style
func IsDotNotation(s string) bool {
	r := regexp.MustCompile(regexCheckDotNotation)
	return r.MatchString(strings.TrimSpace(s))
}

// IsUpperDotNotation detects if a string is upper dot notation style
func IsUpperDotNotation(s string) bool {
	r := regexp.MustCompile(regexCheckUpperDotNotation)
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
