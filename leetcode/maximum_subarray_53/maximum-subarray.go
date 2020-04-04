package maximum_subarray_53

import "math"

// Time complexity: O(n)
func maxSubArray(nums []int) int {
	sum := math.MinInt32
	maxSum := math.MinInt32

	for _, i := range nums {
		if sum+i > i {
			sum += i

		} else {
			sum = i
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
