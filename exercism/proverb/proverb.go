// Package proverb provides function to generate the "For Want of a Nail" proverb.
package proverb

import "fmt"

const (
	line = "For want of a %s the %s was lost."
	last = "And all for the want of a %s."
)

// Proverb generates the "For Want of a Nail" proverb base on input rhyme.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	proverb := make([]string, len(rhyme))
	for i := 0; i < (len(rhyme) - 1); i++ {
		proverb[i] = fmt.Sprintf(line, rhyme[i], rhyme[i+1])
	}
	proverb[len(rhyme)-1] = fmt.Sprintf(last, rhyme[0])

	return proverb
}
