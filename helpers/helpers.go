package helpers

import (
	"regexp"
	"strings"

	"github.com/kennygrant/sanitize"
)

// Compute the length of a string.
func Length(s string) int {
	return len([]rune(s))
}

// Find the minimum between 2 integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Clean a string from special characters, white spaces and
// transform upper case letters to lower case.
func CleanString(s string) string {
	s = strings.Trim(strings.ToLower(s), " ")
	s = sanitize.Accents(s)

	// Remove all other unrecognised characters
	s = regexp.MustCompile(`[^[:alnum:]]`).ReplaceAllString(s, "")

	return s
}
