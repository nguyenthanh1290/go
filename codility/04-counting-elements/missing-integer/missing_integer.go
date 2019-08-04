package missing

import (
	"sort"
)

// O(n*log(n))
// https://app.codility.com/demo/results/trainingPYGD47-NCS/
// Detected time complexity: O(N) or O(N * log(N))
func Solution1(A []int) int {
	sort.Ints(A)
	min := 1
	for _, v := range A {
		if v == min {
			min++
		}
	}

	return min
}

// O(n)
// https://app.codility.com/demo/results/trainingT3JKNV-YMX/
// Detected time complexity: O(N) or O(N * log(N))
func Solution(A []int) int {
	positiveElems := make(map[int]struct{})
	for _, v := range A {
		if v > 0 {
			positiveElems[v] = struct{}{}
		}
	}

	min := 1
	for {
		if _, prs := positiveElems[min]; !prs {
			return min
		}
		min++
	}
}
