package version

// Time complexity: O(logn)
// Space complexity: O(1)
func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2

		if ok := isBadVersion(mid); ok {
			right = mid

		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(version int) bool {
	return true
}
