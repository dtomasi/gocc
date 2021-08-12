package gocc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getAllTestStrings() []string {
	return mergeSlices(
		testStringsSnakeCase,
		testStringsUpperSnakeCase,
		testStringsPascalCase,
		testStringsCamelCase,
		testStringsKebabCase,
		testStringsUpperKebabCase,
	)
}

func checkErr(t *testing.T, result string, matches bool, message string) {
	if !matches {
		t.Errorf("conversion error: %s\n", message)
	}
}

func TestSConverter_ToSnakeCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsSnakeCase(C(testString).ToSnakeCase()))
	}
}

func TestSConverter_ToUpperSnakeCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsUpperSnakeCase(C(testString).ToUpperSnakeCase()))
	}
}

func TestSConverter_ToPascalCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsPascalCase(C(testString).ToPascalCase()))
	}
}

func TestSConverter_ToCamelCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsCamelCase(C(testString).ToCamelCase()))
	}
}

func TestSConverter_ToKebabCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsKebabCase(C(testString).ToKebabCase()))
	}
}

func TestSConverter_ToUpperKebabCase(t *testing.T) {
	for _, testString := range getAllTestStrings() {
		assert.True(t, IsUpperKebabCase(C(testString).ToUpperKebabCase()))
	}
}
