package casestyle_test

import (
	"fmt"
	"github.com/dtomasi/gocc/internal/testutils"
	"github.com/dtomasi/gocc/internal/utils"
	"github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDetectCaseStyle(t *testing.T) {
	for caseStyle, testString := range testutils.TestStyleMapping {
		style, err := casestyle.DetectCaseStyle(testString)
		assert.NoError(t, err)
		assert.Equal(t,
			caseStyle,
			style,
			fmt.Sprintf("Case Style: %s | test string %s", caseStyle, testString),
		)
	}
}

func TestIsSnakeCase(t *testing.T) {
	for _, testString := range testutils.TestStringsSnakeCase {
		assert.True(t, casestyle.IsSnakeCase(testString))
	}

	invalidStrings := utils.MergeSlices(
		testutils.TestStringsUpperSnakeCase,
		testutils.TestStringsPascalCase,
		testutils.TestStringsCamelCase,
		testutils.TestStringsKebabCase,
		testutils.TestStringsUpperKebabCase,
		testutils.TestStringsDotNotation,
		testutils.TestStringsUpperDotNotation,
	)

	for _, testString := range invalidStrings {
		assert.False(t, casestyle.IsSnakeCase(testString), testString)
	}
}

func TestIsPascalCase(t *testing.T) {
	for _, testString := range testutils.TestStringsPascalCase {
		assert.True(t, casestyle.IsPascalCase(testString), testString)
	}

	invalidStrings := utils.MergeSlices(
		testutils.TestStringsSnakeCase,
		testutils.TestStringsUpperSnakeCase,
		testutils.TestStringsCamelCase,
		testutils.TestStringsKebabCase,
		testutils.TestStringsUpperKebabCase,
		testutils.TestStringsDotNotation,
		testutils.TestStringsUpperDotNotation,
	)
	for _, testString := range invalidStrings {
		assert.False(t, casestyle.IsPascalCase(testString), testString)
	}
}

func TestIsCamelCase(t *testing.T) {
	for _, testString := range testutils.TestStringsCamelCase {
		assert.True(t, casestyle.IsCamelCase(testString), testString)
	}

	invalidStrings := utils.MergeSlices(
		testutils.TestStringsSnakeCase,
		testutils.TestStringsUpperSnakeCase,
		testutils.TestStringsPascalCase,
		testutils.TestStringsKebabCase,
		testutils.TestStringsUpperKebabCase,
		testutils.TestStringsDotNotation,
		testutils.TestStringsUpperDotNotation,
	)

	for _, testString := range invalidStrings {
		assert.False(t, casestyle.IsCamelCase(testString), testString)
	}
}

func TestIsKebabCase(t *testing.T) {
	for _, testString := range testutils.TestStringsKebabCase {
		assert.True(t, casestyle.IsKebabCase(testString), testString)
	}

	invalidStrings := utils.MergeSlices(
		testutils.TestStringsSnakeCase,
		testutils.TestStringsUpperSnakeCase,
		testutils.TestStringsPascalCase,
		testutils.TestStringsCamelCase,
		testutils.TestStringsUpperKebabCase,
		testutils.TestStringsDotNotation,
		testutils.TestStringsUpperDotNotation,
	)

	for _, testString := range invalidStrings {
		assert.False(t, casestyle.IsKebabCase(testString), testString)
	}
}

func TestIsDotNotation(t *testing.T) {
	for _, testString := range testutils.TestStringsDotNotation {
		assert.True(t, casestyle.IsDotNotation(testString), testString)
	}

	invalidStrings := utils.MergeSlices(
		testutils.TestStringsSnakeCase,
		testutils.TestStringsUpperSnakeCase,
		testutils.TestStringsPascalCase,
		testutils.TestStringsCamelCase,
		testutils.TestStringsKebabCase,
		testutils.TestStringsUpperKebabCase,
		testutils.TestStringsUpperDotNotation,
	)

	for _, testString := range invalidStrings {
		assert.False(t, casestyle.IsDotNotation(testString), testString)
	}
}

func TestIsCustomDelimiter(t *testing.T) {
	templateString := "custom delimiter test"
	for _, delimiter := range testutils.TestCustomDelimiters {
		testString := strings.ReplaceAll(templateString, " ", delimiter)
		assert.True(t,
			casestyle.IsCustomDelimiter(testString, delimiter),
			fmt.Sprintf("delimiter: %s | test string %s", delimiter, testString),
		)
	}

	invalidStrings := utils.MergeSlices(
		testutils.TestStringsSnakeCase,
		testutils.TestStringsUpperSnakeCase,
		testutils.TestStringsPascalCase,
		testutils.TestStringsCamelCase,
		testutils.TestStringsKebabCase,
		testutils.TestStringsUpperKebabCase,
		testutils.TestStringsDotNotation,
		testutils.TestStringsUpperDotNotation,
	)

	for _, delimiter := range testutils.TestCustomDelimiters {
		for _, testString := range invalidStrings {
			assert.False(t,
				casestyle.IsCustomDelimiter(testString, delimiter),
				fmt.Sprintf("delimiter: %s | test string %s", delimiter, testString),
			)
		}
	}
}

func TestIsUpperCustomDelimiter(t *testing.T) {
	templateString := "CUSTOM DELIMITER TEST"
	for _, delimiter := range testutils.TestCustomDelimiters {
		testString := strings.ReplaceAll(templateString, " ", delimiter)
		assert.True(t,
			casestyle.IsUpperCustomDelimiter(testString, delimiter),
			fmt.Sprintf("delimiter: %s | test string %s", delimiter, testString),
		)
	}

	invalidStrings := utils.MergeSlices(
		testutils.TestStringsSnakeCase,
		testutils.TestStringsUpperSnakeCase,
		testutils.TestStringsPascalCase,
		testutils.TestStringsCamelCase,
		testutils.TestStringsKebabCase,
		testutils.TestStringsUpperKebabCase,
		testutils.TestStringsDotNotation,
		testutils.TestStringsUpperDotNotation,
	)

	for _, delimiter := range testutils.TestCustomDelimiters {
		for _, testString := range invalidStrings {
			assert.False(t,
				casestyle.IsUpperCustomDelimiter(testString, delimiter),
				fmt.Sprintf("delimiter: %s | test string %s", delimiter, testString),
			)
		}
	}
}

func Benchmark_detectCaseStyle_SnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = casestyle.DetectCaseStyle(testutils.TestStringsSnakeCase[0])
	}
}

func Benchmark_detectCaseStyle_KebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = casestyle.DetectCaseStyle(testutils.TestStringsKebabCase[0])
	}
}

func Benchmark_detectCaseStyle_CamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = casestyle.DetectCaseStyle(testutils.TestStringsCamelCase[0])
	}
}

func Benchmark_detectCaseStyle_PascalCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = casestyle.DetectCaseStyle(testutils.TestStringsPascalCase[0])
	}
}

func Benchmark_detectCaseStyle_DotNotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = casestyle.DetectCaseStyle(testutils.TestStringsDotNotation[0])
	}
}
