package gocc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCustomDelimiters = []string{`|`, `*`, `+`, `?`, `^`, `$`, `/`, `\`}

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
		assert.True(t, IsSnakeCase(C(testString).ToSnakeCase()), testString)
	}
}

func BenchmarkConverter_PascalCase_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C(testStringsPascalCase[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_PascalCase_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C(testStringsPascalCase[1]).Convert(StylePascalCase, StyleSnakeCase)
	}
}

func BenchmarkConverter_CamelCase_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C(testStringsCamelCase[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_CamelCase_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C(testStringsCamelCase[1]).Convert(StyleCamelCase, StyleSnakeCase)
	}
}

func BenchmarkConverter_KebabCase_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C(testStringsKebabCase[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_KebabCase_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C(testStringsKebabCase[1]).Convert(StyleKebabCase, StyleSnakeCase)
	}
}


func BenchmarkConverter_DotNotation_ToSnakeCase_with_style_detection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C(testStringsDotNotation[1]).ToSnakeCase()
	}
}

func BenchmarkConverter_DotNotation_ToSnakeCase_with_known_style(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C(testStringsDotNotation[1]).Convert(StyleDotNotation, StyleSnakeCase)
	}
}

func TestSConverter_ToPascalCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsPascalCase(C(testString).ToPascalCase()), testString)
	}
}

func TestSConverter_ToCamelCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsCamelCase(C(testString).ToCamelCase()), testString)
	}
}

func TestSConverter_ToKebabCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsKebabCase(C(testString).ToKebabCase()), testString)
	}
}

func TestSConverter_ToDotNotation(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsDotNotation(C(testString).ToDotNotation()), testString)
	}
}

func TestConverter_ToCustomDelimiter(t *testing.T) {

	for _, delimiter := range testCustomDelimiters {
		for _, testString := range getAllTestStrings() {
			testResultString := C(testString).ToCustomDelimiter(delimiter)
			assert.True(t,
				IsCustomDelimiter(
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
			testResultString := C(testString).ToUpperCustomDelimiter(delimiter)
			assert.True(t,
				IsUpperCustomDelimiter(
					testResultString,
					delimiter,
				),
				fmt.Sprintf("delimiter: %s | test string: %s | test result string: %s", delimiter, testString, testResultString),
			)
		}
	}
}
