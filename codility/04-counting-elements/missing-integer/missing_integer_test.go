package missing

import (
	"testing"
)

var testCases = []struct {
	description string
	input       []int
	expected    int
}{
	{
		description: "positve with duplication",
		input:       []int{1, 3, 6, 4, 1, 2},
		expected:    5,
	},
	{
		description: "positve incresing integers starting from 1",
		input:       []int{1, 2, 3},
		expected:    4,
	},
	{
		description: "negative only",
		input:       []int{-1, -3},
		expected:    1,
	},
	{
		description: "positve and negative",
		input:       []int{-1, -3, 0, 4, 2},
		expected:    1,
	},
	{
		description: "minimal and maximal values",
		input:       []int{-1000000, 1000000},
		expected:    1,
	},
	{
		description: "a single element",
		input:       []int{1},
		expected:    2,
	},
	{
		description: "a single element",
		input:       []int{3},
		expected:    1,
	},
	{
		description: "a single element",
		input:       []int{-2},
		expected:    1,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range testCases {
		output := Solution(c.input)
		if output != c.expected {
			t.Fatalf("FAIL: %v\nInput: %v\nexpected: %v\ngot: %v", c.description, c.input, c.expected, output)
		}

		t.Logf("PASS: %v", c.input)
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			Solution(c.input)
		}
	}
}
