package utils

import (
	"regexp"
	"strings"
)

// UcFirst is a shortcut for strings.Title
// capitalizes the first char.
func UcFirst(s string) string {
	return strings.Title(s)
}

// lcFirst lowers the first char.
func LcFirst(s string) string {
	r := regexp.MustCompile(`^(\b[A-Z])`)

	return r.ReplaceAllStringFunc(s, strings.ToLower)
}

func MergeSlices(args ...[]string) []string {
	mergedSlice := make([]string, 0)
	for _, oneSlice := range args {
		mergedSlice = append(mergedSlice, oneSlice...)
	}

	return mergedSlice
}
