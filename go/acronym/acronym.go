// Package acronym should have a package comment that summarizes what it's about.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	strs := regexp.MustCompile(`[[:alpha:]^']+`).FindAllString(s, -1)
	re := ""

	for _, word := range strs {
		first := word[0]
		re += strings.ToUpper(string(first))
	}

	return re
}
