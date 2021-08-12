package gocc

import (
	"strings"
)

// C returns a new Converter instance
func C(s string) *Converter {
	return &Converter{
		s: s,
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

// ToSnakeCase converts the string value of Converter to a snake case string
func (c *Converter) ToSnakeCase() string {
	return toSnakeCase(c.s)
}

// ToUpperSnakeCase converts the string value of Converter to a upper snake case string
func (c *Converter) ToUpperSnakeCase() string {
	return strings.ToUpper(toSnakeCase(c.s))
}

// ToPascalCase converts the string value of Converter to a pascal case string
func (c *Converter) ToPascalCase() string {
	return toPascalCase(c.s)
}

// ToCamelCase converts the string value of Converter to a camel case string
func (c *Converter) ToCamelCase() string {
	return toCamelCase(c.s)
}

// ToKebabCase converts the string value of Converter to a kebab case string
func (c *Converter) ToKebabCase() string {
	return toKebabCase(c.s)
}

// ToUpperKebabCase converts the string value of Converter to a upper kebab case string
func (c *Converter) ToUpperKebabCase() string {
	return strings.ToUpper(toKebabCase(c.s))
}

// ToDotNotation converts the string value of Converter to a kebab case string
func (c *Converter) ToDotNotation() string {
	return toDotNotation(c.s)
}

// ToUpperDotNotation converts the string value of Converter to a upper kebab case string
func (c *Converter) ToUpperDotNotation() string {
	return strings.ToUpper(toDotNotation(c.s))
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
