package main

import (
	"fmt"
)

// https://app.codility.com/demo/results/trainingR3MP65-A8J/
// Detected time complexity: O(N)
func Solution(H []int) int {
	stones := 0
	stack := []int{}

	for _, h := range H {
		for 0 != len(stack) && stack[len(stack)-1] > h {
			stack = stack[:len(stack)-1]
		}
		if 0 != len(stack) && stack[len(stack)-1] == h {
			continue
		} else {
			stack = append(stack, h)
			stones++
		}
	}

	return stones
}

func main() {
	fmt.Println(Solution([]int{8, 8, 5, 7, 9, 8, 7, 4, 8})) // 7
	fmt.Println(Solution([]int{3, 2, 1}))                   // 3
}
