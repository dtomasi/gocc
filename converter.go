package gocc

import (
	"github.com/dtomasi/gocc/internal/converter"
	cs "github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/pkg/errors"
	"strings"
)

// C returns a new Converter instance.
func C(s string) *Converter {
	return &Converter{
		s: strings.TrimSpace(s),
	}
}

// Converter is a simple struct that holds a string.
type Converter struct {
	s string
}

// Style returns a CaseStyle (int) see: style.go.
func (c *Converter) Style() (cs.CaseStyle, error) {
	return cs.DetectCaseStyle(c.s) // nolint:wrapcheck
}

// Convert converts from -> to given types. This function is more performant than the case/format detection
// on other convert functions.
func (c *Converter) Convert(from cs.CaseStyle, to cs.CaseStyle) (string, error) {
	switch to {
	case cs.StylePascalCase:
		return converter.ToPascalCaseFromCaseStyle(from, c.s) // nolint:wrapcheck
	case cs.StyleCamelCase:
		return converter.ToCamelCaseFromCaseStyle(from, c.s) // nolint:wrapcheck
	case cs.StyleSnakeCase:
		return converter.ToSnakeCaseFromCaseStyle(from, c.s) // nolint:wrapcheck
	case cs.StyleKebabCase:
		return converter.ToKebabCaseFromCaseStyle(from, c.s) // nolint:wrapcheck
	case cs.StyleDotNotation:
		return converter.ToDotNotationFromCaseStyle(from, c.s) // nolint:wrapcheck
	}

	return c.s, errors.Errorf(
		"cannot convert %s to %s. Parameter to is a unknown case style. Got %d",
		c.s,
		from.String(),
		to,
	)
}

func (c *Converter) convertToWithDetect(to cs.CaseStyle) (string, error) {
	caseStyle, err := c.Style()
	if err != nil {
		return c.s, err
	}

	return c.Convert(caseStyle, to)
}

// ConvertCustomDelimiter converts given case style to custom delimiter separated format.
func (c *Converter) ConvertCustomDelimiter(from cs.CaseStyle, delimiter string) (string, error) {
	return converter.ToCustomDelimiterFromCaseStyle(from, c.s, delimiter) // nolint:wrapcheck
}

// ToSnakeCase converts the string value of Converter to a snake case string.
func (c *Converter) ToSnakeCase() (string, error) {
	return c.convertToWithDetect(cs.StyleSnakeCase)
}

// ToPascalCase converts the string value of Converter to a pascal case string.
func (c *Converter) ToPascalCase() (string, error) {
	return c.convertToWithDetect(cs.StylePascalCase)
}

// ToCamelCase converts the string value of Converter to a camel case string.
func (c *Converter) ToCamelCase() (string, error) {
	return c.convertToWithDetect(cs.StyleCamelCase)
}

// ToKebabCase converts the string value of Converter to a kebab case string.
func (c *Converter) ToKebabCase() (string, error) {
	return c.convertToWithDetect(cs.StyleKebabCase)
}

// nolint:godox,nolintlint // ToDotNotation converts the string value of Converter to a kebab case string.
func (c *Converter) ToDotNotation() (string, error) {
	return c.convertToWithDetect(cs.StyleDotNotation)
}

// ToCustomDelimiter converts the string to a string separated by given delimiter.
func (c *Converter) ToCustomDelimiter(delimiter string) (string, error) {
	style, err := c.Style()
	if err != nil {
		return c.s, err
	}

	return c.ConvertCustomDelimiter(style, delimiter)
}

// ToUpperCustomDelimiter converts the string to a string separated by given delimiter in upper case.
func (c *Converter) ToUpperCustomDelimiter(delimiter string) (string, error) {
	style, err := c.Style()
	if err != nil {
		return c.s, err
	}

	result, err := c.ConvertCustomDelimiter(style, delimiter)
	if err != nil {
		return c.s, err
	}

	return strings.ToUpper(result), nil
}

// IsSnakeCase detects if a string is snake case style.
func (c *Converter) IsSnakeCase() bool {
	return cs.IsSnakeCase(c.s)
}

// IsPascalCase detects if a string is pascal case style.
func (c *Converter) IsPascalCase() bool {
	return cs.IsPascalCase(c.s)
}

// IsCamelCase detects if a string is camel case style.
func (c *Converter) IsCamelCase() bool {
	return cs.IsCamelCase(c.s)
}

// IsKebabCase detects if a string is kebab case style.
func (c *Converter) IsKebabCase() bool {
	return cs.IsKebabCase(c.s)
}

// IsDotNotation detects if a string is dot notation style.
func (c *Converter) IsDotNotation() bool {
	return cs.IsDotNotation(c.s)
}

// IsCustomDelimiter detects if a string is separated by a given custom delimiter.
func (c *Converter) IsCustomDelimiter(delimiter string) bool {
	return cs.IsCustomDelimiter(c.s, delimiter)
}

// IsUpperCustomDelimiter detects if a string is separated by a given custom delimiter as upper case string.
func (c *Converter) IsUpperCustomDelimiter(delimiter string) bool {
	return cs.IsUpperCustomDelimiter(c.s, delimiter)
}
