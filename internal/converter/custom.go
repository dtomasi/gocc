package converter

import (
	"github.com/dtomasi/gocc/internal/utils"
	"github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

func ToCustomDelimiterFromCaseStyle(from casestyle.CaseStyle, s string, delimiter string) (string, error) {
	switch from {
	case casestyle.StylePascalCase, casestyle.StyleCamelCase:
		return PascalOrCamelToCustomDelimiter(s, delimiter), nil
	case casestyle.StyleSnakeCase:
		return strings.ToLower(strings.ReplaceAll(s, utils.DelimiterSnakeCase, delimiter)), nil
	case casestyle.StyleKebabCase:
		return strings.ToLower(strings.ReplaceAll(s, utils.DelimiterKebabCase, delimiter)), nil
	case casestyle.StyleDotNotation:
		return strings.ToLower(strings.ReplaceAll(s, utils.DelimiterDotNotation, delimiter)), nil
	}

	return s, errors.Errorf(
		"cannot convert %s to custom delimiter %s. Parameter from is a unknown case style. Got %d",
		s,
		delimiter,
		from,
	)
}

func PascalOrCamelToCustomDelimiter(s string, delimiter string) string {
	s = utils.UcFirst(s)
	r := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	s = r.ReplaceAllStringFunc(s, func(s string) string {
		return strings.Join(strings.Split(s, ""), delimiter)
	})

	return strings.ToLower(s)
}
