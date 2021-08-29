package utils_test

import (
	"github.com/dtomasi/gocc/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils_Ucfirst(t *testing.T) {
	assert.Equal(t, "Ucfirst", utils.UcFirst("ucfirst"))
}

func TestUtils_Lcfirst(t *testing.T) {
	assert.Equal(t, "lcfirst", utils.LcFirst("Lcfirst"))
}
