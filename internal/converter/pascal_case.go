package converter

import (
	"fmt"
	"github.com/dtomasi/gocc/internal/utils"
	"github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

func ToPascalCaseFromCaseStyle(from casestyle.CaseStyle, s string) (string, error) {
	switch from {
	case casestyle.StylePascalCase:
		return s, nil
	case casestyle.StyleCamelCase:
		return utils.UcFirst(s), nil
	case casestyle.StyleSnakeCase:
		return CustomDelimiterToPascalCase(s, utils.DelimiterSnakeCase), nil
	case casestyle.StyleKebabCase:
		return CustomDelimiterToPascalCase(s, utils.DelimiterKebabCase), nil
	case casestyle.StyleDotNotation:
		return CustomDelimiterToPascalCase(s, utils.DelimiterDotNotation), nil
	}

	return s, errors.Errorf(
		"cannot convert %s to %s. Parameter from is a unknown case style. Got %d",
		s,
		casestyle.StylePascalCase.String(),
		from,
	)
}

func CustomDelimiterToPascalCase(s string, delimiter string) string {
	regexDelimiter := delimiter
	if utils.DelimiterIsReservedChar(delimiter) {
		regexDelimiter = `\` + delimiter
	}

	regexString := fmt.Sprintf(`(^[A-Za-z])|%s([A-Za-z])`, regexDelimiter)

	var r = regexp.MustCompile(regexString)

	return r.ReplaceAllStringFunc(strings.ToLower(s), func(s string) string {
		return strings.ToUpper(strings.ReplaceAll(s, delimiter, ""))
	})
}
