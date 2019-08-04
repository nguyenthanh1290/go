package div

import (
	"testing"
)

func TestSolution(t *testing.T) {
	for _, c := range testCases {
		output := Solution(c.input[0], c.input[1], c.input[2])
		if output != c.expected {
			t.Fatalf("\nFAIL: %v\nInput: [%v..%v] and %v\nexpected: %v\ngot: %v", c.description, c.input[0], c.input[1], c.input[2], c.expected, output)
		}

		t.Logf("PASS: %v", c.description)
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			Solution(c.input[0], c.input[1], c.input[2])
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
		input:       []int{6, 11, 2},
		expected:    3,
	},
	{
		description: "No.2",
		input:       []int{11, 14, 2},
		expected:    2,
	},
	{
		description: "No.3",
		input:       []int{11, 345, 17},
		expected:    20,
	},
	{
		description: "No.4",
		input:       []int{0, 1, 11},
		expected:    1,
	},
	{
		description: "No.5",
		input:       []int{10, 10, 5},
		expected:    1,
	},
	{
		description: "No.6",
		input:       []int{10, 10, 7},
		expected:    0,
	},
	{
		description: "No.7",
		input:       []int{10, 10, 20},
		expected:    0,
	},
}
