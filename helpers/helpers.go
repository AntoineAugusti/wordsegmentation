package helpers

import (
	"regexp"
	"strings"

	"github.com/kennygrant/sanitize"
)

func Length(s string) int {
	return len([]rune(s))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CleanString(s string) string {
	s = strings.Trim(strings.ToLower(s), " ")
	s = sanitize.Accents(s)

	// Replace certain joining characters with a dash
	s = regexp.MustCompile(`[ &_=+:]`).ReplaceAllString(s, "-")

	// Remove all other unrecognised characters
	s = regexp.MustCompile(`[^[:alnum:]-]`).ReplaceAllString(s, "")

	// Remove any multiple dashes caused by replacements above
	s = regexp.MustCompile(`[\-]+`).ReplaceAllString(s, "-")

	return s
}
