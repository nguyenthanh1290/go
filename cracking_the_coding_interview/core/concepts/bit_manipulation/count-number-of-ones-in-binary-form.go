package bitmnp

// Count the number of ones in the binary representation of the given number.
// Time complexity: O(K), where K is the number of ones present in the binary form of the given number.
// Space complexity: O(1)

// Have x.
// (x-1) will have all the bits same as x, except for the rightmost 1 in x and all the bits to the right of the rightmost 1.
// x & (x-1) will have all the bits equal to the x except for the rightmost 1 in x.
// x & (x-1) = 6 & 5 = (110)2 & (101)2 = (100)2
func countOnes(x int) int {
	count := 0

	for x != 0 {
		x = x & (x - 1)
		count++
	}

	return count
}
