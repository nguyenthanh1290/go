// Package acronym provides a function that converts a pharse to its ancronym.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate converts a pharse to its ancronym.
func Abbreviate(s string) string {
	reg := regexp.MustCompile(`\b[a-zA-Z]`)
	abbr := reg.FindAllString(s, -1)

	return strings.ToUpper(strings.Join(abbr, ""))
}
