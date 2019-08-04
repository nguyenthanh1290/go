package bucket

// SortInts sorts the input slice using bucket algorithm.
// Time complexity: O(n)
func SortInts(a []int) []int {
	min := 1<<63 - 1
	max := -1 << 63
	for _, v := range a {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	bucket := make([]int, (max-min)+1)
	for _, v := range a {
		idx := v - min
		bucket[idx]++
	}

	sorted := make([]int, len(a))
	idx := 0
	for i, v := range bucket {
		if v > 0 {
			for j := 0; j < v; j++ {
				sorted[idx] = i + min
				idx++
			}
		}
	}

	return sorted
}
