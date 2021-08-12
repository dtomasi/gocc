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
)

var styleStringMap = []string{
	"unknown",
	"snake",
	"upper_snake",
	"pascal",
	"camel",
	"kebab",
	"upper_kebab",
}

func (cs CaseStyle) String() string {
	return styleStringMap[cs]
}
