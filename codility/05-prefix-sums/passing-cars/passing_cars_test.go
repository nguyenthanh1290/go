package passingcars

import (
	"testing"
)

func TestPassingCars(t *testing.T) {
	for _, c := range testCases {
		output := PassingCars(c.input)
		if output != c.expected {
			t.Fatalf("FAIL: %v\nInput: %v\nexpected: %v\ngot: %v", c.description, c.input, c.expected, output)
		}

		t.Logf("PASS: %v", c.input)
	}
}

func BenchmarkPassingCars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			PassingCars(c.input)
		}
	}
}

var testCases = []struct {
	description string
	input       []int
	expected    int
}{
	{
		description: "five pairs of passing cars",
		input:       []int{0, 1, 0, 1, 1},
		expected:    5,
	},
	{
		description: "five pairs of passing cars",
		input:       []int{0, 1, 0, 1, 1, 0},
		expected:    5,
	},
	{
		description: "seven pairs of passing cars",
		input:       []int{0, 1, 0, 1, 1, 1},
		expected:    7,
	},
	{
		description: "eight pairs of passing cars",
		input:       []int{0, 1, 0, 1, 1, 0, 1},
		expected:    8,
	},
	{
		description: "zero pairs of passing cars",
		input:       []int{1, 1, 1},
		expected:    0,
	},
	{
		description: "zero pairs of passing cars",
		input:       []int{0, 0, 0},
		expected:    0,
	},
}
