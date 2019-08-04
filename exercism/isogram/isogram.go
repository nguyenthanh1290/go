// Package isogram provides functions that determines if a word or phrase is an isogram.
package isogram

import "unicode"

// IsIsogram returns true if input is an isogram, else returns false.
func IsIsogram(input string) bool {
	dict := map[rune]int{}

	for _, c := range input {
		if c == ' ' || c == '-' {
			continue
		}
		if _, prs := dict[unicode.ToUpper(c)]; prs {
			return false
		}
		dict[unicode.ToUpper(c)]++
	}
	return true
}
