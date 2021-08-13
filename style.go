package gocc

type CaseStyle int

const (
	StyleUnknown CaseStyle = iota
	StylePascalCase
	StyleCamelCase
	StyleSnakeCase
	StyleKebabCase
	StyleDotNotation
	StyleUpperSnakeCase
	StyleUpperKebabCase
	StyleUpperDotNotation
)

var styleStringMap = []string{
	"unknown",
	"pascal",
	"camel",
	"snake",
	"kebab",
	"dot_notation",
	"upper_snake",
	"upper_kebab",
	"upper_dot_notation",
}

func (cs CaseStyle) String() string {
	return styleStringMap[cs]
}

func (cs CaseStyle) IsCaseSensitiveStyle() bool {
	return cs > 0 && cs <= 2
}

func (cs CaseStyle) IsLowerStyle() bool {
	return cs <= 5 && cs > 2
}

func (cs CaseStyle) IsUpperStyle() bool {
	return cs > 5
}
