package gocc

import (
	"bytes"
	"strings"
)

// ucFirst is a shortcut for strings.Title
// capitalizes the first char
func ucFirst(s string) string {
	return strings.Title(s)
}

// lcFirst lowers the first char
func lcFirst(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
