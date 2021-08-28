package gocc_test

import (
	"fmt"
	"github.com/dtomasi/gocc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCustomDelimiters = []string{`|`, `*`, `+`, `?`, `^`, `$`, `/`, `\`} //nolint:gochecknoglobals

func getAllTestStrings() []string {
	return mergeSlices(
		testStringsSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsKebabCase,
		testStringsDotNotation,
	)
}

func TestSConverter_ToSnakeCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, gocc.IsSnakeCase(gocc.C(testString).ToSnakeCase()), testString)
	}
}

func BenchmarkConverter_PascalCase_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gocc.C(testStringsPascalCase[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_PascalCase_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gocc.C(testStringsPascalCase[1]).Convert(gocc.StylePascalCase, gocc.StyleSnakeCase)
	}
}

func BenchmarkConverter_CamelCase_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gocc.C(testStringsCamelCase[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_CamelCase_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gocc.C(testStringsCamelCase[1]).Convert(gocc.StyleCamelCase, gocc.StyleSnakeCase)
	}
}

func BenchmarkConverter_KebabCase_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gocc.C(testStringsKebabCase[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_KebabCase_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gocc.C(testStringsKebabCase[1]).Convert(gocc.StyleKebabCase, gocc.StyleSnakeCase)
	}
}

func BenchmarkConverter_DotNotation_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gocc.C(testStringsDotNotation[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_DotNotation_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gocc.C(testStringsDotNotation[1]).Convert(gocc.StyleDotNotation, gocc.StyleSnakeCase)
	}
}

func TestSConverter_ToPascalCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, gocc.IsPascalCase(gocc.C(testString).ToPascalCase()), testString)
	}
}

func TestSConverter_ToCamelCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, gocc.IsCamelCase(gocc.C(testString).ToCamelCase()), testString)
	}
}

func TestSConverter_ToKebabCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, gocc.IsKebabCase(gocc.C(testString).ToKebabCase()), testString)
	}
}

func TestSConverter_ToDotNotation(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, gocc.IsDotNotation(gocc.C(testString).ToDotNotation()), testString)
	}
}

func TestConverter_ToCustomDelimiter(t *testing.T) {
	for _, delimiter := range testCustomDelimiters {
		for _, testString := range getAllTestStrings() {
			testResultString := gocc.C(testString).ToCustomDelimiter(delimiter)
			assert.True(t,
				gocc.IsCustomDelimiter(
					testResultString,
					delimiter,
				),
				fmt.Sprintf("delimiter: %s | test string: %s | test result string: %s", delimiter, testString, testResultString),
			)
		}
	}
}

func TestConverter_ToUpperCustomDelimiter(t *testing.T) {
	for _, delimiter := range testCustomDelimiters {
		for _, testString := range getAllTestStrings() {
			testResultString := gocc.C(testString).ToUpperCustomDelimiter(delimiter)
			assert.True(t,
				gocc.IsUpperCustomDelimiter(
					testResultString,
					delimiter,
				),
				fmt.Sprintf("delimiter: %s | test string: %s | test result string: %s", delimiter, testString, testResultString),
			)
		}
	}
}
