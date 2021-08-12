package gocc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils_Ucfirst(t *testing.T) {
	assert.Equal(t, "Ucfirst", ucFirst("ucfirst"))
}

func TestUtils_Lcfirst(t *testing.T) {
	assert.Equal(t, "lcfirst", lcFirst("Lcfirst"))
}
