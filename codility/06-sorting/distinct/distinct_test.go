package distinct

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
		input:       []int{2, 1, 1, 2, 3, 1},
		expected:    3,
	},
	{
		description: "No.2",
		input:       []int{},
		expected:    0,
	},
	{
		description: "No.1",
		input:       []int{2},
		expected:    1,
	},
}
