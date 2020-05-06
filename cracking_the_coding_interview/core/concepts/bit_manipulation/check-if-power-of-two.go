package bitmnp

// How to check if a given number is a power of 2?
// Time complexity: O(1)
// Space complexity: O(1)

// If x is a power of 2 then x & (x-1) will be 0.
/*
Ex:
	x		= 4 = (100)2
	x - 1	= 3 = (011)2
	---
	x & (x-1) = 4 & 3 = (100)2 & (011)2 = (000)2
*/
func isPowerOfTwo(x int) bool {
	return (x != 0) && ((x & (x - 1)) == 0)
}
