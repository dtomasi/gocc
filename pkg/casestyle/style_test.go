package casestyle_test

import (
	"github.com/dtomasi/gocc/pkg/casestyle"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCaseStyle_String(t *testing.T) {
	assert.IsType(t, reflect.String, reflect.ValueOf(casestyle.StylePascalCase.String()).Kind())
	assert.IsType(t, reflect.String, reflect.ValueOf(casestyle.StyleKebabCase.String()).Kind())
	assert.IsType(t, reflect.String, reflect.ValueOf(casestyle.StyleCamelCase.String()).Kind())
	assert.IsType(t, reflect.String, reflect.ValueOf(casestyle.StyleDotNotation.String()).Kind())
	assert.IsType(t, reflect.String, reflect.ValueOf(casestyle.StyleSnakeCase.String()).Kind())
}
