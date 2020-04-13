package array

// Approach #1 Brute Force
// Time complexity : O(n^2). We consider every possible subarray by traversing over the complete array for every start point possible.
// Space complexity : O(1). Only two variables zeroeszeroes and onesones are required.
func findMaxLengthBruteForce(nums []int) int {
	maxLen := 0

	for start := 0; start < len(nums); start++ {
		zeroes := 0
		ones := 0

		for end := start; end < len(nums); end++ {
			if nums[end] == 0 {
				zeroes++

			} else {
				ones++
			}

			if zeroes == ones {
				maxLen = func(a, b int) int {
					if a > b {
						return a
					}
					return b
				}(maxLen, end-start+1)
			}
		}
	}

	return maxLen
}

// Approach #3 Using HashMap
// Time complexity : O(n). The entire array is traversed only once.
// Space complexity : O(n). Maximum size of the HashMap map will be n, if all the elements are either 1 or 0.
func findMaxLength(nums []int) int {
	maxLen := 0
	m := map[int]int{
		0: -1,
	}
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++

		} else {
			count--
		}

		if c, ok := m[count]; ok {
			maxLen = func(a, b int) int {
				if a > b {
					return a
				}
				return b
			}(maxLen, i-c)

		} else {
			m[count] = i
		}
	}

	return maxLen
}
