package triangle

import (
	"sort"
)

// https://app.codility.com/demo/results/training8VR3Q4-TKG/
// Detected time complexity:
// O(N*log(N))
func Solution(A []int) int {
	sort.Ints(A)
	for i := 0; i < len(A)-2; i++ {
		if A[i]+A[i+1] > A[i+2] {
			return 1
		}
	}
	return 0
}
