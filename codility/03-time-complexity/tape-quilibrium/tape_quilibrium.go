package main

import (
	"fmt"
	"math"
)

func Solution(A []int) int {
	min := math.MaxInt32
	total := sum(A...)
	lSum := 0
	for i := 0; i < len(A)-1; i++ {
		lSum += A[i]
		rSum := total - lSum

		diff := abs(lSum - rSum)
		if diff < min {
			min = diff
		}
	}

	return min
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(Solution([]int{3, 1, 2, 4, 3}))
}

// https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/
// https://app.codility.com/demo/results/training5MXBJQ-SRA/
// Detected time complexity:
// O(N)