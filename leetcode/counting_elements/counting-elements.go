package counting_elements

// Time complexity: O(n)
// Space complexity: O(n)
func countElements(arr []int) int {
	bucket := make(map[int]struct{}, len(arr))
	count := 0

	for _, a := range arr {
		bucket[a] = struct{}{}
	}

	for _, a := range arr {
		if _, ok := bucket[a+1]; ok {
			count++
		}
	}

	return count
}
