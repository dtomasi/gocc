package gocc

import "regexp"

// detectCaseStyle returns the detected style as int (CaseStyle)
// see constants.go
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
	}

	return StyleInvalid
}

// IsSnakeCase detects if a string is snake case style
func IsSnakeCase(s string) bool {
	r := regexp.MustCompile(regexCheckSnackCase)
	return r.MatchString(s)
}

// IsUpperSnakeCase detects if a string is snake case style
func IsUpperSnakeCase(s string) bool {
	r := regexp.MustCompile(regexCheckUpperSnakeCase)
	return r.MatchString(s)
}

// IsPascalCase detects if a string is pascal case style
func IsPascalCase(s string) bool {
	r := regexp.MustCompile(regexCheckPascalCase)
	return r.MatchString(s)
}

// IsCamelCase detects if a string is camel case style
func IsCamelCase(s string) bool {
	r := regexp.MustCompile(regexCheckCamelCase)
	return r.MatchString(s)
}

// IsKebabCase detects if a string is kebab case style
func IsKebabCase(s string) bool {
	r := regexp.MustCompile(regexCheckKebabCase)
	return r.MatchString(s)
}

// IsUpperKebabCase detects if a string is kebab case style
func IsUpperKebabCase(s string) bool {
	r := regexp.MustCompile(regexCheckUpperKebabCase)
	return r.MatchString(s)
}

// IsDotNotation detects if a string is dot notation style
func IsDotNotation(s string) bool {
	r := regexp.MustCompile(regexCheckDotNotation)
	return r.MatchString(s)
}

// IsUpperDotNotation detects if a string is upper dot notation style
func IsUpperDotNotation(s string) bool {
	r := regexp.MustCompile(regexCheckUpperDotNotation)
	return r.MatchString(s)
}
