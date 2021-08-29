package converter

import (
	"github.com/dtomasi/gocc/internal/utils"
	"github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/pkg/errors"
)

func ToCamelCaseFromCaseStyle(from casestyle.CaseStyle, s string) (string, error) {
	switch from {
	case casestyle.StyleCamelCase:
		return s, nil
	case casestyle.StylePascalCase:
		return utils.LcFirst(s), nil
	case casestyle.StyleSnakeCase:
		return CustomDelimiterToCamelCase(s, utils.DelimiterSnakeCase), nil
	case casestyle.StyleKebabCase:
		return CustomDelimiterToCamelCase(s, utils.DelimiterKebabCase), nil
	case casestyle.StyleDotNotation:
		return CustomDelimiterToCamelCase(s, utils.DelimiterDotNotation), nil
	}

	return s, errors.Errorf(
		"cannot convert %s to %s. Parameter from is a unknown case style. Got %d",
		s,
		casestyle.StyleCamelCase.String(),
		from,
	)
}

func CustomDelimiterToCamelCase(s string, delimiter string) string {
	return utils.LcFirst(CustomDelimiterToPascalCase(s, delimiter))
}
