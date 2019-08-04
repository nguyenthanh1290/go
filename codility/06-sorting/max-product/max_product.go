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
