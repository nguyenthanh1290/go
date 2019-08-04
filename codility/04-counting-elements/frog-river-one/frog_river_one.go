package main

import (
	"fmt"
)

func Solution(X int, A []int) int {
	expectedTotal := X * (X + 1) / 2
	total := 0
	seen := make(map[int]struct{})

	for i, v := range A {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			total += v
		}
		if total == expectedTotal {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(Solution(5, []int{1, 3, 1, 4, 2, 3, 5, 4})) // expected 6
	fmt.Println(Solution(3, []int{1, 3, 1, 3, 2, 1, 3}))    // expected 4
}

// https://app.codility.com/programmers/lessons/4-counting_elements/frog_river_one/
// https://app.codility.com/demo/results/trainingC7Q4JH-QU6/
// Detected time complexity:
// O(N)