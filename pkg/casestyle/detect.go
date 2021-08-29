package casestyle

import (
	"errors"
	"fmt"
	"github.com/dtomasi/gocc/internal/utils"
	"regexp"
	"strings"
)

var (
	errUnknownStyle = errors.New("detection error")
)

// DetectCaseStyle returns the detected style as int (CaseStyle)
// see style.go.
func DetectCaseStyle(s string) (CaseStyle, error) {
	switch {
	case IsPascalCase(s):
		return StylePascalCase, nil
	case IsCamelCase(s):
		return StyleCamelCase, nil
	}

	delimiter, found := detectDelimitedString(s)
	if !found {
		return -1, fmt.Errorf("%w: %s", errUnknownStyle, s)
	}

	switch delimiter {
	case utils.DelimiterSnakeCase:
		return StyleSnakeCase, nil
	case utils.DelimiterKebabCase:
		return StyleKebabCase, nil
	case utils.DelimiterDotNotation:
		return StyleDotNotation, nil
	}

	return -1, fmt.Errorf("%w: %s", errUnknownStyle, s)
}

func detectDelimitedString(s string) (string, bool) {
	r := regexp.MustCompile(utils.RegexDelimiterDetection)
	subMatch := r.FindStringSubmatch(s)

	if len(subMatch) > 0 {
		// this is a required validation step to check if a string only contains detected delimiter
		if IsCustomDelimiter(s, subMatch[1]) {
			return subMatch[1], true
		}
	}

	return "", false
}

// IsSnakeCase detects if a string is snake case style.
func IsSnakeCase(s string) bool {
	r := regexp.MustCompile(utils.RegexCheckSnackCase)

	return r.MatchString(strings.TrimSpace(s))
}

// IsPascalCase detects if a string is pascal case style.
func IsPascalCase(s string) bool {
	r := regexp.MustCompile(utils.RegexCheckPascalCase)

	return r.MatchString(strings.TrimSpace(s))
}

// IsCamelCase detects if a string is camel case style.
func IsCamelCase(s string) bool {
	r := regexp.MustCompile(utils.RegexCheckCamelCase)

	return r.MatchString(strings.TrimSpace(s))
}

// IsKebabCase detects if a string is kebab case style.
func IsKebabCase(s string) bool {
	r := regexp.MustCompile(utils.RegexCheckKebabCase)

	return r.MatchString(strings.TrimSpace(s))
}

// IsDotNotation detects if a string is dot notation style.
func IsDotNotation(s string) bool {
	r := regexp.MustCompile(utils.RegexCheckDotNotation)

	return r.MatchString(strings.TrimSpace(s))
}

// IsCustomDelimiter detects if a string is separated by a given custom delimiter.
func IsCustomDelimiter(s string, delimiter string) bool {
	r := regexp.MustCompile(utils.BuildCustomDelimiterRegex(delimiter))

	return r.MatchString(strings.TrimSpace(s))
}

// IsUpperCustomDelimiter detects if a string is separated by a given custom delimiter as upper case string.
func IsUpperCustomDelimiter(s string, delimiter string) bool {
	r := regexp.MustCompile(utils.BuildUpperCustomDelimiterRegex(delimiter))

	return r.MatchString(strings.TrimSpace(s))
}
