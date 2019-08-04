package main

import (
	"fmt"
)

// https://app.codility.com/demo/results/trainingRJR4TK-YDV/
// Detected time complexity: O(N)
func max(A []int) int {
	forwardSum := make([]int, len(A))
	backwardSum := make([]int, len(A))

	for i := 1; i < len(A)-1; i++ {
		if forwardSum[i-1]+A[i] > 0 {
			forwardSum[i] = forwardSum[i-1] + A[i]
		} else {
			forwardSum[i] = 0
		}
	}
	for i := len(A) - 2; i >= 1; i-- {
		if backwardSum[i+1]+A[i] > 0 {
			backwardSum[i] = backwardSum[i+1] + A[i]
		} else {
			backwardSum[i] = 0
		}
	}

	max := 0
	for i := 0; i < len(A)-2; i++ {
		if forwardSum[i]+backwardSum[i+2] > max {
			max = forwardSum[i] + backwardSum[i+2]
		}
	}
	return max
}

func main() {
	fmt.Println(max([]int{3, 2, 6, -1, 4, 5, -1, 2}))
	// forwardSum: [0, 2, 8, 7, 11, 16, 15, 0]
	// backwardSum: [0, 16, 14, 8, 9, 5, 0, 0]
	// result: 17
}
