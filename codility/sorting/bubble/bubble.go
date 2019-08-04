package bubble

// SortInts sorts the input slice using pubble algorithm.
// Time complexity:
// - worst case (reverse sorted): O(n^2)
// - best case (already sorted): O(n)
// - average case: O(n^2)
// Space complexity: O(1)
// Stable: Yes
func SortInts(a []int) []int {
	sorted := make([]int, len(a))
	copy(sorted, a)

	for i := 0; i < len(sorted); i++ {
		swapped := false
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return sorted
}
