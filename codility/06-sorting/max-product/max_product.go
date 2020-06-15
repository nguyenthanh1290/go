package product

import "sort"

// https://app.codility.com/demo/results/training3STRZ3-6SX/
// Detected time complexity:
// O(N * log(N))
func Solution(A []int) int {
	sort.Ints(A)
	i := len(A) - 1
	rMax := A[i] * A[i-1] * A[i-2]
	lMax := A[0] * A[1] * A[i]

	if rMax > lMax {
		return rMax
	}
	return lMax
}

// Time complexity: O(n)
// Space complexity: O(1)
func useSingleScan(nums []int) int {
	min1, min2 := 1<<63-1, 1<<63-1
	max1, max2, max3 := -1<<63, -1<<63, -1<<63

	for _, n := range nums {
		if n <= min1 {
			min1, min2 = n, min1
		} else if n <= min2 {
			min2 = n
		}

		if n >= max1 {
			max1, max2, max3 = n, max1, max2
		} else if n >= max2 {
			max2, max3 = n, max2
		} else if n >= max3 {
			max3 = n
		}
	}

	p1 := min1 * min2 * max1
	p2 := max3 * max2 * max1

	if p1 > p2 {
		return p1
	}
	return p2
}
