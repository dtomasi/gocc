package gocc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var (
	testStyleMapping = map[CaseStyle]string{
		StyleUnknown:          "in-va_lid",
		StyleSnakeCase:        "snake_case",
		StylePascalCase:       "PascalCase",
		StyleCamelCase:        "camelCase",
		StyleKebabCase:        "kebab-case",
		StyleDotNotation:      "dot.notation",
	}

	testStringsSnakeCase = []string{
		"snake_case",
		"snake_case_multi_words",
	}
	testStringsUpperSnakeCase = []string{
		"UPPER_SNAKE",
		"UPPER_SNAKE_CASE",
	}
	testStringsPascalCase = []string{
		"PascalCase",
		"PascalCaseMultiWords",
	}
	testStringsCamelCase = []string{
		"camelCase",
		"camelCaseMultiWords",
	}
	testStringsKebabCase = []string{
		"kebab-case",
		"kebab-case-multi-words",
	}
	testStringsUpperKebabCase = []string{
		"UPPER-KEBAB",
		"UPPER-KEBAB-CASE",
	}
	testStringsDotNotation = []string{
		"dot.notation",
		"dot.notation.multi.words",
	}
	testStringsUpperDotNotation = []string{
		"DOT.NOTATION",
		"DOT.NOTATION.MULTI.WORDS",
	}
)

func mergeSlices(args ...[]string) []string {
	mergedSlice := make([]string, 0)
	for _, oneSlice := range args {
		mergedSlice = append(mergedSlice, oneSlice...)
	}

	return mergedSlice
}

func TestDetectCaseStyle(t *testing.T) {

	for caseStyle, testString := range testStyleMapping {
		assert.Equal(t, caseStyle, detectCaseStyle(testString), fmt.Sprintf("Case Style: %s | test string %s", caseStyle, testString))
	}
}

func TestIsSnakeCase(t *testing.T) {

	for _, testString := range testStringsSnakeCase {
		assert.True(t, IsSnakeCase(testString))
	}

	invalidStrings := mergeSlices(
		testStringsUpperSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsKebabCase,
		testStringsUpperKebabCase,
		testStringsDotNotation,
		testStringsUpperDotNotation,
	)

	for _, testString := range invalidStrings {
		assert.False(t, IsSnakeCase(testString), testString)
	}
}

func TestIsPascalCase(t *testing.T) {

	for _, testString := range testStringsPascalCase {
		assert.True(t, IsPascalCase(testString), testString)
	}

	invalidStrings := mergeSlices(
		testStringsSnakeCase,
		testStringsUpperSnakeCase,
		testStringsCamelCase,
		testStringsKebabCase,
		testStringsUpperKebabCase,
		testStringsDotNotation,
		testStringsUpperDotNotation,
	)
	for _, testString := range invalidStrings {
		assert.False(t, IsPascalCase(testString), testString)
	}
}

func TestIsCamelCase(t *testing.T) {

	for _, testString := range testStringsCamelCase {
		assert.True(t, IsCamelCase(testString), testString)
	}

	invalidStrings := mergeSlices(
		testStringsSnakeCase,
		testStringsUpperSnakeCase,
		testStringsPascalCase,
		testStringsKebabCase,
		testStringsUpperKebabCase,
		testStringsDotNotation,
		testStringsUpperDotNotation,
	)

	for _, testString := range invalidStrings {
		assert.False(t, IsCamelCase(testString), testString)
	}
}

func TestIsKebabCase(t *testing.T) {

	for _, testString := range testStringsKebabCase {
		assert.True(t, IsKebabCase(testString), testString)
	}

	invalidStrings := mergeSlices(
		testStringsSnakeCase,
		testStringsUpperSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsUpperKebabCase,
		testStringsDotNotation,
		testStringsUpperDotNotation,
	)

	for _, testString := range invalidStrings {
		assert.False(t, IsKebabCase(testString), testString)
	}
}

func TestIsDotNotation(t *testing.T) {

	for _, testString := range testStringsDotNotation {
		assert.True(t, IsDotNotation(testString), testString)
	}

	invalidStrings := mergeSlices(
		testStringsSnakeCase,
		testStringsUpperSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsKebabCase,
		testStringsUpperKebabCase,
		testStringsUpperDotNotation,
	)

	for _, testString := range invalidStrings {
		assert.False(t, IsDotNotation(testString), testString)
	}
}

func TestIsCustomDelimiter(t *testing.T) {

	templateString := "custom delimiter test"
	for _, delimiter := range testCustomDelimiters {
		testString := strings.ReplaceAll(templateString, " ", delimiter)
		assert.True(t,
			IsCustomDelimiter(testString, delimiter),
			fmt.Sprintf("delimiter: %s | test string %s", delimiter, testString),
		)
	}

	invalidStrings := mergeSlices(
		testStringsSnakeCase,
		testStringsUpperSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsKebabCase,
		testStringsUpperKebabCase,
		testStringsDotNotation,
		testStringsUpperDotNotation,
	)

	for _, delimiter := range testCustomDelimiters {
		for _, testString := range invalidStrings {
			assert.False(t,
				IsCustomDelimiter(testString, delimiter),
				fmt.Sprintf("delimiter: %s | test string %s", delimiter, testString),
			)
		}
	}
}

func TestIsUpperCustomDelimiter(t *testing.T) {

	templateString := "CUSTOM DELIMITER TEST"
	for _, delimiter := range testCustomDelimiters {
		testString := strings.ReplaceAll(templateString, " ", delimiter)
		assert.True(t,
			IsUpperCustomDelimiter(testString, delimiter),
			fmt.Sprintf("delimiter: %s | test string %s", delimiter, testString),
		)
	}

	invalidStrings := mergeSlices(
		testStringsSnakeCase,
		testStringsUpperSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsKebabCase,
		testStringsUpperKebabCase,
		testStringsDotNotation,
		testStringsUpperDotNotation,
	)

	for _, delimiter := range testCustomDelimiters {
		for _, testString := range invalidStrings {
			assert.False(t,
				IsUpperCustomDelimiter(testString, delimiter),
				fmt.Sprintf("delimiter: %s | test string %s", delimiter, testString),
			)
		}
	}
}

func Benchmark_detectCaseStyle_SnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectCaseStyle(testStringsSnakeCase[0])
	}
}

func Benchmark_detectCaseStyle_KebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectCaseStyle(testStringsKebabCase[0])
	}
}

func Benchmark_detectCaseStyle_CamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectCaseStyle(testStringsCamelCase[0])
	}
}

func Benchmark_detectCaseStyle_PascalCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectCaseStyle(testStringsPascalCase[0])
	}
}

func Benchmark_detectCaseStyle_DotNotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		detectCaseStyle(testStringsDotNotation[0])
	}
}
