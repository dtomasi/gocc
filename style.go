package gocc

type CaseStyle int

const (
	StyleUnknown CaseStyle = iota
	StylePascalCase
	StyleCamelCase
	StyleSnakeCase
	StyleKebabCase
	StyleDotNotation
)

var styleStringMap = []string{
	"unknown",
	"pascal",
	"camel",
	"snake",
	"kebab",
	"dot_notation",
}

func (cs CaseStyle) String() string {
	return styleStringMap[cs]
}
