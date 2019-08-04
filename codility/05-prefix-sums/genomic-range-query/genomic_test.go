package genomic

import (
	"testing"
)

func TestSolution(t *testing.T) {
	for _, c := range testCases {
		output := Solution(c.dna, c.p, c.q)
		if len(output) != len(c.expected) {
			t.Fatalf("FAIL:\nDNA: %v\nQuery: %v %v\nexpected: %v\ngot: %v", c.dna, c.p, c.q, c.expected, output)
		}
		for i, v := range output {
			if v != c.expected[i] {
				t.Fatalf("FAIL:\nDNA: %v\nQuery: %v %v\nexpected: %v\ngot: %v", c.dna, c.p, c.q, c.expected, output)
			}
		}

		t.Logf("PASS: %v", c.dna)
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			Solution(c.dna, c.p, c.q)
		}
	}
}

var testCases = []struct {
	dna      string
	p        []int
	q        []int
	expected []int
}{
	{
		dna:      "CAGCCTA",
		p:        []int{2, 5, 0},
		q:        []int{4, 5, 6},
		expected: []int{2, 4, 1},
	},
	{
		dna:      "C",
		p:        []int{0},
		q:        []int{0},
		expected: []int{2},
	},
	{
		dna:      "CA",
		p:        []int{0},
		q:        []int{1},
		expected: []int{1},
	},
	{
		dna:      "CC",
		p:        []int{0, 0, 1},
		q:        []int{1, 0, 1},
		expected: []int{2, 2, 2},
	},
	{
		dna:      "AAAC",
		p:        []int{0, 0, 1, 3, 2},
		q:        []int{1, 2, 3, 3, 3},
		expected: []int{1, 1, 1, 2, 1},
	},
}
