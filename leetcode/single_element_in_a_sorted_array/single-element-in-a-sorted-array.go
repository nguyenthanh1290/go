package single

// Time complexity: O(logN)
// Space complexity: O(1)
func singleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if mid&1 == 0 { // mid is even
			if nums[mid] != nums[mid+1] {
				right = mid

			} else {
				left = mid + 2
			}

		} else { // mid is odd
			if nums[mid] != nums[mid-1] {
				right = mid - 1

			} else {
				left = mid + 1
			}
		}
	}

	return nums[left]
}
