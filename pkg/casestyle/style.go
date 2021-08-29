package casestyle

type CaseStyle int

const (
	StylePascalCase CaseStyle = iota
	StyleCamelCase
	StyleSnakeCase
	StyleKebabCase
	StyleDotNotation
)

var styleStringMap = []string{
	"pascal",
	"camel",
	"snake",
	"kebab",
	"dot_notation",
}

func (cs CaseStyle) String() string {
	if int(cs) > len(styleStringMap) || int(cs) < 0 {
		return "UNKNOWN"
	}

	return styleStringMap[cs]
}
