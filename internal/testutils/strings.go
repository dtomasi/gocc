package testutils

import "github.com/dtomasi/gocc/pkg/casestyle"

var (
	TestCustomDelimiters = []string{`|`, `*`, `+`, `?`, `^`, `$`, `/`, `\`} //nolint:gochecknoglobals

	TestStyleMapping = map[casestyle.CaseStyle]string{
		casestyle.StyleSnakeCase:   "snake_case",
		casestyle.StylePascalCase:  "PascalCase",
		casestyle.StyleCamelCase:   "camelCase",
		casestyle.StyleKebabCase:   "kebab-case",
		casestyle.StyleDotNotation: "dot.notation",
	}

	TestStringsSnakeCase = []string{
		"snake_case",
		"snake_case_multi_words",
	}
	TestStringsUpperSnakeCase = []string{
		"UPPER_SNAKE",
		"UPPER_SNAKE_CASE",
	}
	TestStringsPascalCase = []string{
		"PascalCase",
		"PascalCaseMultiWords",
	}
	TestStringsCamelCase = []string{
		"camelCase",
		"camelCaseMultiWords",
	}
	TestStringsKebabCase = []string{
		"kebab-case",
		"kebab-case-multi-words",
	}
	TestStringsUpperKebabCase = []string{
		"UPPER-KEBAB",
		"UPPER-KEBAB-CASE",
	}
	TestStringsDotNotation = []string{
		"dot.notation",
		"dot.notation.multi.words",
	}
	TestStringsUpperDotNotation = []string{
		"DOT.NOTATION",
		"DOT.NOTATION.MULTI.WORDS",
	}
)
