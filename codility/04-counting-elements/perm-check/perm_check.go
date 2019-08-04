package main

import (
	"fmt"
)

func Solution(A []int) int {
	counter := make(map[int]struct{})
	for _, v := range A {
		if v > len(A) {
			return 0
		}
		counter[v] = struct{}{}
	}
	if len(counter) == len(A) {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(Solution([]int{4, 1, 3, 2})) // 1
	fmt.Println(Solution([]int{4, 1, 3})) // 0
}

// https://app.codility.com/demo/results/training9WRW9M-F43/
// Detected time complexity:
// O(N) or O(N * log(N))