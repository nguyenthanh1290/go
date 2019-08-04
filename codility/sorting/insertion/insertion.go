package insertion

// SortInts sorts the input slice using insertion algorithm.
// Time complexity:
// - worst case (reverse sorted): O(n^2)
// - best case (already sorted): O(n)
// - average case: O(n^2)
// Space complexity: O(1)
// Stable: Yes
// Usage: insertion sort runs in linear time on nearly sorted data.
func SortInts(a []int) []int {
	sorted := make([]int, len(a))
	copy(sorted, a)

	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j] > key {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}

	return sorted
}
