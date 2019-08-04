package main

import (
	"fmt"
)

func Solution(A []int) int {
	if len(A) == 0 {
		return 1
	}

	sumA := 0
	for _, v := range A {
		sumA += v
	}

	n := len(A) + 1
	sumM := n * (n + 1) / 2

	return sumM - sumA
}

func main() {
	fmt.Println(Solution([]int{2, 3, 1, 5}))
}

// https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/
// https://app.codility.com/demo/results/trainingJSDCQ6-8XK/
// Detected time complexity:
// O(N)