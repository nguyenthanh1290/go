package brackets

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
	input       string
	expected    int
}{
	{
		description: "No.1",
		input:       "",
		expected:    1,
	},
	{
		description: "No.2",
		input:       "{[()()]}",
		expected:    1,
	},
	{
		description: "No.3",
		input:       "([)()]",
		expected:    0,
	},
	{
		description: "No.4",
		input:       "()[]",
		expected:    1,
	},
	{
		description: "No.5",
		input:       "()[{}]",
		expected:    1,
	},
	{
		description: "No.6",
		input:       ")(",
		expected:    0,
	},
	{
		description: "No.6",
		input:       ")",
		expected:    0,
	},
}
