package gocc

import (
	"regexp"
	"strings"
)

// ucFirst is a shortcut for strings.Title
// capitalizes the first char
func ucFirst(s string) string {
	return strings.Title(s)
}

// lcFirst lowers the first char
func lcFirst(s string) string {
	r := regexp.MustCompile(`^(\b[A-Z])`)
	return r.ReplaceAllStringFunc(s, func(s string) string {
		return strings.ToLower(s)
	})
}
