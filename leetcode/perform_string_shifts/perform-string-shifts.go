package string

// Time complexity: O(n)
// Space complexity: O(n)
func stringShift(s string, shift [][]int) string {
	left := 0

	for _, s := range shift {
		if s[0] == 0 {
			left += s[1]

		} else {
			left -= s[1]
		}
	}

	left %= len(s)

	if left == 0 {
		return s

	} else if left > 0 { // perform left shift
		return s[left:] + s[:left]

	} else { // perform right shift
		right := -left

		return s[len(s)-right:] + s[:len(s)-right]
	}
}
