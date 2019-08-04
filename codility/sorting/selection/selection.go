package selection

// SortInts sorts the input slice using selection algorithm.
// Time complexity:
// - worst case: O(n^2)
// - best case:
// - average case: O(n^2)
// Space complexity: O(1)
// Stable: Yes
func SortInts(a []int) []int {
	sorted := make([]int, len(a))
	copy(sorted, a)

	for i := 0; i < len(sorted)-1; i++ {
		idxOfMinElem := i
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j] < sorted[idxOfMinElem] {
				idxOfMinElem = j
			}
		}
		sorted[i], sorted[idxOfMinElem] = sorted[idxOfMinElem], sorted[i]
	}

	return sorted
}
