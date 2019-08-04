package distinct

import "sort"

// https://app.codility.com/demo/results/trainingJ2JGAF-HCV/
// Detected time complexity:
// O(N*log(N)) or O(N)
func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)
	result := 1
	for i := 1; i < len(A); i++ {
		if A[i] != A[i-1] {
			result++
		}
	}
	return result
}

// https://app.codility.com/demo/results/trainingTEP555-WWU/
// Detected time complexity:
// O(N)
func Solution1(A []int) int {
	m := make(map[int]struct{})
	for _, v := range A {
		m[v] = struct{}{}
	}

	return len(m)
}

// by counting elements
// Notice that we have to know the range of the sorted values.
// If all the elements are in the set {0, 1, . . . , k},
// then the array used for counting should be of size k + 1

// The time complexity here is O(n + k).
// We need additional memory O(k) to count all the elements
func SolutionXXX(A []int, K int) int {
	counter := make([]int, K + 1) //init all elements of the slice to 0 value
	for _, v := range A {
		counter[v]++
	}
	result := 0
	for _, v := range counter {
		if v > 0 {
			result++
		}
	}
	return result
}