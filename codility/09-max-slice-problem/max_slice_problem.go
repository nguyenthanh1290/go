package main

import (
	"fmt"
)

// prefix sums
// O(n**2)
func max1(A []int) int {
	prefixSums := make([]int, len(A)+1)
	for i := 1; i < len(A)+1; i++ {
		prefixSums[i] = prefixSums[i-1] + A[i-1]
	}

	maxSum := 0
	for p := 0; p < len(A); p++ {
		for q := p; q < len(A); q++ {
			sum := prefixSums[q+1] - prefixSums[p]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

// without prefix sums
// O(n**2)
func max2(A []int) int {
	maxSum := 0
	for p := 0; p < len(A); p++ {
		sum := 0
		for q := p; q < len(A); q++ {
			sum += A[q]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

// O(n)
func max(A []int) int {
	maxSum := 0
	sum := 0

	for _, a := range A {
		if sum+a > 0 {
			sum += a

		} else {
			sum = 0
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func main() {
	fmt.Println(max([]int{1, 2, 3, 4, 5}))          // 15
	fmt.Println(max([]int{5, -7, 3, 5, -2, 4, -1})) // 10
	fmt.Println(max([]int{5, -1}))                  // 5
}
