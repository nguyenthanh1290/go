package triangle

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
		input:       []int{10, 2, 5, 1, 8, 20},
		expected:    1,
	},
	{
		description: "No.2",
		input:       []int{10, 50, 5, 1},
		expected:    0,
	},
	{
		description: "No.2",
		input:       []int{1, 2147483647, 2147483647},
		expected:    1,
	},
}
