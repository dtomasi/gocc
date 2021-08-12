package gocc

type CaseStyle int

const (
	StyleInvalid CaseStyle = iota
	StyleSnakeCase
	StyleUpperSnakeCase
	StylePascalCase
	StyleCamelCase
	StyleKebabCase
	StyleUpperKebabCase
	StyleDotNotation
	StyleUpperDotNotation
)

var styleStringMap = []string{
	"unknown",
	"snake",
	"upper_snake",
	"pascal",
	"camel",
	"kebab",
	"upper_kebab",
	"dot_notation",
	"upper_dot_notation",
}

func (cs CaseStyle) String() string {
	return styleStringMap[cs]
}
