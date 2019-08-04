package main

import (
	"fmt"
)

// https://app.codility.com/demo/results/trainingPESXMD-AA7/
// Detected time complexity: O(N)
func max(A []int) int {
	maxSlice := A[0]
	maxEnding := A[0]

	for i := 1; i < len(A); i++ {
		if A[i] > maxEnding+A[i] {
			maxEnding = A[i]
		} else {
			maxEnding += A[i]
		}
		if maxEnding > maxSlice {
			maxSlice = maxEnding
		}
	}

	return maxSlice
}

func main() {
	fmt.Println(max([]int{3, 2, -6, 4, 0})) // 5
	fmt.Println(max([]int{-10}))            // -10
	fmt.Println(max([]int{-10, -3}))        // -3
	fmt.Println(max([]int{5, 1}))           // 6
	fmt.Println(max([]int{-2, 1}))          // 1
	fmt.Println(max([]int{1, 3}))           // 4
}
