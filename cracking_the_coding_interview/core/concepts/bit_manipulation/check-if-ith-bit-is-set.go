package bitmnp

// Check if the ith bit is set in the binary form of the given number.
// Time complexity: O(1)
// Space complexity: O(1)
func isSet(n int, i uint) bool {
	return (n & (1 << i)) != 0
}
