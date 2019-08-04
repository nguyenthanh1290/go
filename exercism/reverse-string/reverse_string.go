package reverse

import (
	"strings"
)

// String reverses the input.
func String(s string) string {
	var builder strings.Builder
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}
	return builder.String()
}
