package gocc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testStyleMapping = map[CaseStyle]string{
		StyleInvalid:        "in-va_lid",
		StyleSnakeCase:      "snake_case",
		StyleUpperSnakeCase: "UPPER_SNAKE_CASE",
		StylePascalCase:     "PascalCase",
		StyleCamelCase:      "camelCase",
		StyleKebabCase:      "kebab-case",
		StyleUpperKebabCase: "UPPER-KEBAB-CASE",
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
		assert.Equal(t, caseStyle, detectCaseStyle(testString))
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
	)

	for _, testString := range invalidStrings {
		assert.False(t, IsSnakeCase(testString), testString)
	}
}

func TestIsUpperSnakeCase(t *testing.T) {

	for _, testString := range testStringsUpperSnakeCase {
		assert.True(t, IsUpperSnakeCase(testString), testString)
	}

	invalidStrings := mergeSlices(
		testStringsSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsKebabCase,
		testStringsUpperKebabCase,
	)

	for _, testString := range invalidStrings {
		assert.False(t, IsUpperSnakeCase(testString), testString)
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
	)

	for _, testString := range invalidStrings {
		assert.False(t, IsKebabCase(testString), testString)
	}
}

func TestIsUpperKebabCase(t *testing.T) {

	for _, testString := range testStringsUpperKebabCase {
		assert.True(t, IsUpperKebabCase(testString), testString)
	}

	invalidStrings := mergeSlices(
		testStringsSnakeCase,
		testStringsUpperSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsKebabCase,
	)

	for _, testString := range invalidStrings {
		assert.False(t, IsUpperKebabCase(testString), testString)
	}
}
