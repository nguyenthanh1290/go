// Package bob provides a function that simulates Bob's responses to people's questions.
package bob

import (
	"strings"
)

// Remark is a convenience type for identifying what kind of remark it is.
type Remark string

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(string(r), "?")
}

func (r Remark) isYelling() bool {
	allNonLetter := true
	for _, c := range r {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			allNonLetter = false
		}
		if c >= 'a' && c <= 'z' {
			return false
		}
	}
	return !allNonLetter
}

func (r Remark) isExasperated() bool {
	return r.isYelling() && r.isQuestion()
}

func (r Remark) isSilence() bool {
	return r == ""
}

// Hey simulates Bob's reponse to inputted remark.
func Hey(remark string) string {
	r := Remark(strings.TrimSpace(remark))
	switch {
	case r.isExasperated():
		return "Calm down, I know what I'm doing!"
	case r.isQuestion():
		return "Sure."
	case r.isYelling():
		return "Whoa, chill out!"
	case r.isSilence():
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
