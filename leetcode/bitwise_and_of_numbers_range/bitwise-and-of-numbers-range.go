package bitwise

// Time complexity: O(1)
// Space complexity: O(1)
func rangeBitwiseAnd(m int, n int) int {
	var count uint

	for m != n {
		m >>= 1
		n >>= 1

		count++
	}

	return m << count
}
