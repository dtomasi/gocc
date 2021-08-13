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
		testStringsUpperSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsKebabCase,
		testStringsUpperKebabCase,
		testStringsDotNotation,
		testStringsUpperDotNotation,
	)
}

func TestSConverter_ToSnakeCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsSnakeCase(C(testString).ToSnakeCase()), testString)
	}
}

func TestSConverter_ToUpperSnakeCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsUpperSnakeCase(C(testString).ToUpperSnakeCase()), testString)
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

func TestSConverter_ToUpperKebabCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsUpperKebabCase(C(testString).ToUpperKebabCase()), testString)
	}
}

func TestSConverter_ToDotNotation(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsDotNotation(C(testString).ToDotNotation()), testString)
	}
}

func TestSConverter_ToUpperDotNotation(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsUpperDotNotation(C(testString).ToUpperDotNotation()), testString)
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
			assert.True(t,
				IsUpperCustomDelimiter(
					C(testString).ToUpperCustomDelimiter(delimiter),
					delimiter,
				),
				fmt.Sprintf("delimiter: %s | test string %s", delimiter, testString),
			)
		}
	}
}
