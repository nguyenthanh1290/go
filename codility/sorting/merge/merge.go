package merge

// SortInts sorts the input slice using merge algorithm.
// Time complexity:
// - worst case (reverse sorted): O(n*logn)
// - best case (already sorted): O(n*logn)
// - average case: O(n*logn)
// Space complexity: O(n)
// Stable: Yes
// Usage: .
func SortInts(a []int) []int {
	sorted := make([]int, len(a))
	copy(sorted, a)

	sort(sorted, 0, len(sorted)-1)
	return sorted
}

func sort(a []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		sort(a, l, m)
		sort(a, m+1, r)
		merge(a, l, m, r)
	}
}

func merge(a []int, l, m, r int) {
	L := make([]int, m-l+1)
	R := make([]int, r-m)
	copy(L, a[l:m+1])
	copy(R, a[m+1:r+1])

	i := 0
	j := 0
	k := l
	for i < len(L) && j < len(R) {
		if L[i] <= R[j] {
			a[k] = L[i]
			i++
		} else {
			a[k] = R[j]
			j++
		}
		k++
	}

	for i < len(L) {
		a[k] = L[i]
		i++
		k++
	}

	for j < len(R) {
		a[k] = R[j]
		j++
		k++
	}
}
