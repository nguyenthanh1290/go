// Package strand provides function to transcibe dna to rna
package strand

import (
	"bytes"
)

// ToRNA transcibes dna to rna
func ToRNA(dna string) string {
	if dna == "" {
		return ""
	}

	var buf bytes.Buffer
	for _, nucleotide := range dna {
		switch nucleotide {
		case 'G':
			buf.WriteRune('C')
		case 'C':
			buf.WriteRune('G')
		case 'T':
			buf.WriteRune('A')
		case 'A':
			buf.WriteRune('U')
		}
	}
	return buf.String()
}
