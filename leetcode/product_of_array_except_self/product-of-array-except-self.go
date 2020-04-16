package array

// Approach 1: Left and Right product lists
// Time complexity : O(N)
// Space complexity : O(N)
func productExceptSelfLR(nums []int) []int {
	lProducts := make([]int, len(nums))
	lProducts[0] = 1
	for i := 1; i < len(nums); i++ {
		lProducts[i] = lProducts[i-1] * nums[i-1]
	}

	rProducts := make([]int, len(nums))
	rProducts[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rProducts[i] = rProducts[i+1] * nums[i+1]
	}

	products := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		products[i] = lProducts[i] * rProducts[i]
	}

	return products
}

// Approach 2: O(1) space approach
// Time complexity : O(N)
// Space complexity : O(1)
func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))
	
	products[0] = 1
	for i := 1; i < len(nums); i++ {
		products[i] = products[i-1] * nums[i-1]
	}

	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		products[i] *= r
		r *= nums[i]
	}

	return products
}
