package handler

import (
	"strings"
)

// low returns the given string with all letters in lower-case.
func low(s string) string {
	return strings.ToLower(s)
}
