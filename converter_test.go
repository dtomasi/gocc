package gocc_test

import (
	"fmt"
	"github.com/dtomasi/gocc"
	"github.com/dtomasi/gocc/internal/testutils"
	"github.com/dtomasi/gocc/internal/utils"
	"github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func getAllTestStrings() []string {
	return utils.MergeSlices(
		testutils.TestStringsSnakeCase,
		testutils.TestStringsPascalCase,
		testutils.TestStringsCamelCase,
		testutils.TestStringsKebabCase,
		testutils.TestStringsDotNotation,
	)
}

func TestConverter_ToNoMatch(t *testing.T) {
	_, err := gocc.C("something:that_is|not.valid").Convert(-1, -1)
	assert.Error(t, err)
}

func TestSConverter_ToSnakeCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		result, err := gocc.C(testString).ToSnakeCase()
		assert.NoError(t, err)
		assert.True(t, casestyle.IsSnakeCase(result), testString)
	}
}

func BenchmarkConverter_PascalCase_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gocc.C(testutils.TestStringsPascalCase[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_PascalCase_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gocc.C(testutils.TestStringsPascalCase[1]).Convert(casestyle.StylePascalCase, casestyle.StyleSnakeCase)
	}
}

func BenchmarkConverter_CamelCase_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gocc.C(testutils.TestStringsCamelCase[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_CamelCase_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gocc.C(testutils.TestStringsCamelCase[1]).Convert(casestyle.StyleCamelCase, casestyle.StyleSnakeCase)
	}
}

func BenchmarkConverter_KebabCase_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gocc.C(testutils.TestStringsKebabCase[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_KebabCase_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gocc.C(testutils.TestStringsKebabCase[1]).Convert(casestyle.StyleKebabCase, casestyle.StyleSnakeCase)
	}
}

func BenchmarkConverter_DotNotation_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gocc.C(testutils.TestStringsDotNotation[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_DotNotation_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gocc.C(testutils.TestStringsDotNotation[1]).Convert(casestyle.StyleDotNotation, casestyle.StyleSnakeCase)
	}
}

func TestSConverter_ToPascalCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		result, err := gocc.C(testString).ToPascalCase()
		assert.NoError(t, err)
		assert.True(t, casestyle.IsPascalCase(result), testString)
	}
}

func TestSConverter_ToCamelCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		result, err := gocc.C(testString).ToCamelCase()
		assert.NoError(t, err)
		assert.True(t, casestyle.IsCamelCase(result), testString)
	}
}

func TestSConverter_ToKebabCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		result, err := gocc.C(testString).ToKebabCase()
		assert.NoError(t, err)
		assert.True(t, casestyle.IsKebabCase(result), testString)
	}
}

func TestSConverter_ToDotNotation(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		result, err := gocc.C(testString).ToDotNotation()
		assert.NoError(t, err)
		assert.True(t, casestyle.IsDotNotation(result), testString)
	}
}

func TestConverter_ToCustomDelimiter(t *testing.T) {
	for _, delimiter := range testutils.TestCustomDelimiters {
		for _, testString := range getAllTestStrings() {
			result, err := gocc.C(testString).ToCustomDelimiter(delimiter)
			assert.NoError(t, err)
			assert.True(t,
				casestyle.IsCustomDelimiter(
					result,
					delimiter,
				),
				fmt.Sprintf("delimiter: %s | test string: %s | test result string: %s", delimiter, testString, result),
			)
		}
	}
}

func TestConverter_ToUpperCustomDelimiter(t *testing.T) {
	for _, delimiter := range testutils.TestCustomDelimiters {
		for _, testString := range getAllTestStrings() {
			result, err := gocc.C(testString).ToUpperCustomDelimiter(delimiter)
			assert.NoError(t, err)
			assert.True(t,
				casestyle.IsUpperCustomDelimiter(
					result,
					delimiter,
				),
				fmt.Sprintf("delimiter: %s | test string: %s | test result string: %s", delimiter, testString, result),
			)
		}
	}
}

func TestConverter_IsSnakeCase(t *testing.T) {
	for _, testString := range testutils.TestStringsSnakeCase {
		assert.True(t, gocc.C(testString).IsSnakeCase())
	}
}

func TestConverter_IsCamelCase(t *testing.T) {
	for _, testString := range testutils.TestStringsCamelCase {
		assert.True(t, gocc.C(testString).IsCamelCase())
	}
}

func TestConverter_IsDotNotation(t *testing.T) {
	for _, testString := range testutils.TestStringsDotNotation {
		assert.True(t, gocc.C(testString).IsDotNotation())
	}
}

func TestConverter_IsPascalCase(t *testing.T) {
	for _, testString := range testutils.TestStringsPascalCase {
		assert.True(t, gocc.C(testString).IsPascalCase())
	}
}

func TestConverter_IsKebabCase(t *testing.T) {
	for _, testString := range testutils.TestStringsKebabCase {
		assert.True(t, gocc.C(testString).IsKebabCase())
	}
}

func TestConverter_IsCustomDelimiter(t *testing.T) {
	stringTmpl := "test string template"
	for _, testDelimiter := range testutils.TestCustomDelimiters {
		testString := strings.ReplaceAll(stringTmpl, " ", testDelimiter)
		assert.True(t, gocc.C(testString).IsCustomDelimiter(testDelimiter))
	}
}

func TestConverter_IsUpperCustomDelimiter(t *testing.T) {
	stringTmpl := "TEST STRING TEMPLATE"
	for _, testDelimiter := range testutils.TestCustomDelimiters {
		testString := strings.ReplaceAll(stringTmpl, " ", testDelimiter)
		assert.True(t, gocc.C(testString).IsUpperCustomDelimiter(testDelimiter))
	}
}
