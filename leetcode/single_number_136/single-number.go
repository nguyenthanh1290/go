package single_number_136

// Bit Manipulation
// Algorithm:
/*
	If we take XOR of zero and some bit, it will return that bit
		a ^ 0 = a
	If we take XOR of two same bits, it will return 0
		a ^ a = 0
	a ^ b ^ a = (a ^ a) ^ b = 0 ^ b = b

	So we can XOR all bits together to find the unique number.
*/
// Complexity:
/*
	- Time complexity : O(n)
	- Space complexity : O(1)
*/
func singleNumberUsingXOR(nums []int) int {
	a := 0
	for i := 0; i < len(nums); i++ {
		a ^= nums[i]
	}
	return a
}

// Map
// Algorithm:
/*
	1. Iterate through all elements in nums and set up key/value pair.
	2. Return the element which appeared only once.
*/
// Complexity:
/*
	- Time complexity : O(n)
	- Space complexity : O(n)
*/
func singleNumberUsingMap(nums []int) int {
	count := make(map[int]int, len(nums))

	for _, n := range nums {
		count[n]++
	}

	for k, v := range count {
		if v == 1 {
			return k
		}
	}

	return 0
}

// Math
// Algorithm: 2∗(a+b+c)−(a+a+b+b+c) = c
// Complexity:
/*
	- Time complexity : O(n)
	- Space complexity : O(n)
*/
func singleNumberUsingMath(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	sumOfSet := 0
	sumOfNums := 0

	for _, n := range nums {
		sumOfNums += n

		set[n] = struct{}{}
	}

	for k := range set {
		sumOfSet += k
	}

	return 2*sumOfSet - sumOfNums
}
