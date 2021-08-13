package gocc

import (
	"strings"
)

// C returns a new Converter instance
func C(s string) *Converter {
	return &Converter{
		s: strings.TrimSpace(s),
	}
}

// Converter is a simple struct that holds a string
type Converter struct {
	s string
}

// Style returns a CaseStyle (int) see: style.go
func (c *Converter) Style() CaseStyle {
	return detectCaseStyle(c.s)
}

// Convert converts from -> to given types. This function is more performant than the case/format detection
// on other convert functions
func (c *Converter) Convert(from CaseStyle, to CaseStyle) string {
	switch to {
	case StylePascalCase:
		return toPascalCaseFromCaseStyle(from, c.s)
	case StyleCamelCase:
		return toCamelCaseFromCaseStyle(from, c.s)
	case StyleSnakeCase:
		return toSnakeCaseFromCaseStyle(from, c.s)
	case StyleUpperSnakeCase:
		return strings.ToUpper(toSnakeCaseFromCaseStyle(from, c.s))
	case StyleKebabCase:
		return toKebabCaseFromCaseStyle(from, c.s)
	case StyleUpperKebabCase:
		return strings.ToUpper(toKebabCaseFromCaseStyle(from, c.s))
	case StyleDotNotation:
		return toDotNotationFromCaseStyle(from, c.s)
	case StyleUpperDotNotation:
		return strings.ToUpper(toDotNotationFromCaseStyle(from, c.s))
	}

	return c.s
}

// ConvertCustomDelimiter converts given case style to custom delimiter separated format
func (c *Converter) ConvertCustomDelimiter(from CaseStyle, delimiter string) string {
	return toCustomDelimiterFromCaseStyle(from, c.s, delimiter)
}

// ToSnakeCase converts the string value of Converter to a snake case string
func (c *Converter) ToSnakeCase() string {
	return c.Convert(c.Style(), StyleSnakeCase)
}

// ToUpperSnakeCase converts the string value of Converter to a upper snake case string
func (c *Converter) ToUpperSnakeCase() string {
	return c.Convert(c.Style(), StyleUpperSnakeCase)
}

// ToPascalCase converts the string value of Converter to a pascal case string
func (c *Converter) ToPascalCase() string {
	return c.Convert(c.Style(), StylePascalCase)
}

// ToCamelCase converts the string value of Converter to a camel case string
func (c *Converter) ToCamelCase() string {
	return c.Convert(c.Style(), StyleCamelCase)
}

// ToKebabCase converts the string value of Converter to a kebab case string
func (c *Converter) ToKebabCase() string {
	return c.Convert(c.Style(), StyleKebabCase)
}

// ToUpperKebabCase converts the string value of Converter to a upper kebab case string
func (c *Converter) ToUpperKebabCase() string {
	return c.Convert(c.Style(), StyleUpperKebabCase)
}

// ToDotNotation converts the string value of Converter to a kebab case string
func (c *Converter) ToDotNotation() string {
	return c.Convert(c.Style(), StyleDotNotation)
}

// ToUpperDotNotation converts the string value of Converter to a upper kebab case string
func (c *Converter) ToUpperDotNotation() string {
	return c.Convert(c.Style(), StyleUpperDotNotation)
}

// ToCustomDelimiter converts the string to a string separated by given delimiter
func (c *Converter) ToCustomDelimiter(delimiter string) string {
	return c.ConvertCustomDelimiter(c.Style(), delimiter)
}

// ToUpperCustomDelimiter converts the string to a string separated by given delimiter in upper case
func (c *Converter) ToUpperCustomDelimiter(delimiter string) string {
	return strings.ToUpper(c.ConvertCustomDelimiter(c.Style(), delimiter))
}

// IsSnakeCase detects if a string is snake case style
func (c *Converter) IsSnakeCase() bool {
	return IsSnakeCase(c.s)
}

// IsUpperSnakeCase detects if a string is snake case style
func (c *Converter) IsUpperSnakeCase() bool {
	return IsUpperSnakeCase(c.s)
}

// IsPascalCase detects if a string is pascal case style
func (c *Converter) IsPascalCase() bool {
	return IsPascalCase(c.s)
}

// IsCamelCase detects if a string is camel case style
func (c *Converter) IsCamelCase() bool {
	return IsCamelCase(c.s)
}

// IsKebabCase detects if a string is kebab case style
func (c *Converter) IsKebabCase() bool {
	return IsKebabCase(c.s)
}

// IsUpperKebabCase detects if a string is kebab case style
func (c *Converter) IsUpperKebabCase() bool {
	return IsUpperKebabCase(c.s)
}

// IsDotNotation detects if a string is dot notation style
func (c *Converter) IsDotNotation() bool {
	return IsDotNotation(c.s)
}

// IsUpperDotNotation detects if a string is upper dot notation style
func (c *Converter) IsUpperDotNotation() bool {
	return IsUpperDotNotation(c.s)
}

// IsCustomDelimiter detects if a string is separated by a given custom delimiter
func (c *Converter) IsCustomDelimiter(delimiter string) bool {
	return IsCustomDelimiter(c.s, delimiter)
}

// IsUpperCustomDelimiter detects if a string is separated by a given custom delimiter as upper case string
func (c *Converter) IsUpperCustomDelimiter(delimiter string) bool {
	return IsUpperCustomDelimiter(c.s, delimiter)
}
