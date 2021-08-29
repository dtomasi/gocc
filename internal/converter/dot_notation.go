package converter

import (
	"github.com/dtomasi/gocc/internal/utils"
	"github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/pkg/errors"
	"strings"
)

func ToDotNotationFromCaseStyle(from casestyle.CaseStyle, s string) (string, error) {
	switch from {
	case casestyle.StyleDotNotation:
		return s, nil
	case casestyle.StylePascalCase, casestyle.StyleCamelCase:
		return PascalOrCamelCaseToDotNotation(s), nil
	case casestyle.StyleKebabCase:
		return KebabCaseToDotNotation(s), nil
	case casestyle.StyleSnakeCase:
		return SnakeCaseToDotNotation(s), nil
	}

	return s, errors.Errorf(
		"cannot convert %s to %s. Parameter from is a unknown case style. Got %d",
		s,
		casestyle.StyleDotNotation.String(),
		from,
	)
}

func PascalOrCamelCaseToDotNotation(s string) string {
	return PascalOrCamelToCustomDelimiter(s, utils.DelimiterDotNotation)
}

func KebabCaseToDotNotation(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, utils.DelimiterKebabCase, utils.DelimiterDotNotation))
}

func SnakeCaseToDotNotation(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, utils.DelimiterSnakeCase, utils.DelimiterDotNotation))
}
