package converter

import (
	"github.com/dtomasi/gocc/internal/utils"
	"github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/pkg/errors"
	"strings"
)

func ToSnakeCaseFromCaseStyle(from casestyle.CaseStyle, s string) (string, error) {
	switch from {
	case casestyle.StyleSnakeCase:
		return s, nil
	case casestyle.StylePascalCase, casestyle.StyleCamelCase:
		return PascalOrCamelCaseToSnakeCase(s), nil
	case casestyle.StyleKebabCase:
		return KebabCaseToSnakeCase(s), nil
	case casestyle.StyleDotNotation:
		return DotNotationToSnakeCase(s), nil
	}

	return s, errors.Errorf(
		"cannot convert %s to %s. Parameter from is a unknown case style. Got %d",
		s,
		casestyle.StyleSnakeCase.String(),
		from,
	)
}

func PascalOrCamelCaseToSnakeCase(s string) string {
	return PascalOrCamelToCustomDelimiter(s, utils.DelimiterSnakeCase)
}

func KebabCaseToSnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, utils.DelimiterKebabCase, utils.DelimiterSnakeCase))
}

func DotNotationToSnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, utils.DelimiterDotNotation, utils.DelimiterSnakeCase))
}
