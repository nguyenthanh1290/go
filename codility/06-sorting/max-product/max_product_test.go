package product

import (
	"testing"
)

func TestSolution(t *testing.T) {
	for _, c := range testCases {
		output := Solution(c.input)
		if output != c.expected {
			t.Fatalf("\nFAIL: %v\nInput: %v\nexpected: %v\ngot: %v", c.description, c.input, c.expected, output)
		}

		t.Logf("PASS: %v", c.description)
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			Solution(c.input)
		}
	}
}

var testCases = []struct {
	description string
	input       []int
	expected    int
}{
	{
		description: "No.1",
		input:       []int{-3, 1, 2, -2, 5, 6},
		expected:    60,
	},
	{
		description: "No.1",
		input:       []int{-5, 5, -5, 4},
		expected:    125,
	},
}
