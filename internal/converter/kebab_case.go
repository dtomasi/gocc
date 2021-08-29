package converter

import (
	"github.com/dtomasi/gocc/internal/utils"
	"github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/pkg/errors"
	"strings"
)

func ToKebabCaseFromCaseStyle(from casestyle.CaseStyle, s string) (string, error) {
	switch from {
	case casestyle.StyleKebabCase:
		return s, nil
	case casestyle.StylePascalCase, casestyle.StyleCamelCase:
		return PascalOrCamelCaseToKebabCase(s), nil
	case casestyle.StyleSnakeCase:
		return SnakeCaseToKebabCase(s), nil
	case casestyle.StyleDotNotation:
		return DotNotationToKebabCase(s), nil
	}

	return s, errors.Errorf(
		"cannot convert %s to %s. Parameter from is a unknown case style. Got %d",
		s,
		casestyle.StyleKebabCase.String(),
		from,
	)
}

func PascalOrCamelCaseToKebabCase(s string) string {
	return PascalOrCamelToCustomDelimiter(s, utils.DelimiterKebabCase)
}

func SnakeCaseToKebabCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, utils.DelimiterSnakeCase, utils.DelimiterKebabCase))
}

func DotNotationToKebabCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, utils.DelimiterDotNotation, utils.DelimiterKebabCase))
}
