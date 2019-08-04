package minavg

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
		description: "Positive integers",
		input:       []int{4, 2, 2, 5, 1, 5, 8},
		expected:    1,
	},
	{
		description: "Nagative integers",
		input:       []int{-3, -5, -8, -4, -10},
		expected:    2,
	},
}
