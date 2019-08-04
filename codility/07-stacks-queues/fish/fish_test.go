package fish

import (
	"testing"
)

func TestSolution(t *testing.T) {
	for _, c := range testCases {
		output := Solution(c.a, c.b)
		if output != c.expected {
			t.Fatalf("\nFAIL: %v\nA: %v\nB: %v\nexpected: %v\ngot: %v", c.description, c.a, c.b, c.expected, output)
		}

		t.Logf("PASS: %v", c.description)
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			Solution(c.a, c.b)
		}
	}
}

var testCases = []struct {
	description string
	a           []int
	b           []int
	expected    int
}{
	{
		description: "No.1",
		a:           []int{4, 3, 2, 1, 5},
		b:           []int{0, 1, 0, 0, 0},
		expected:    2,
	},
	{
		description: "No.3",
		a:           []int{4, 3, 4, 2, 1, 5},
		b:           []int{0, 1, 1, 0, 0, 0},
		expected:    2,
	},
	{
		description: "No.2",
		a:           []int{4, 3, 4, 2, 5, 1},
		b:           []int{0, 1, 1, 0, 0, 1},
		expected:    3,
	},
}
