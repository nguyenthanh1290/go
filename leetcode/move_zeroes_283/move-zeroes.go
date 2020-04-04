package move_zeroes_283

// Space Complexity : O(1)
// Time Complexity: O(n^2)
func moveZeroes1(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
			}
		}
	}
}

// Space Complexity : O(1)
// Time Complexity: O(n)
func moveZeroes2(nums []int) {
	zeroPtr := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroPtr], nums[i] = nums[i], nums[zeroPtr]

			zeroPtr++
		}
	}
}

// Space Complexity : O(1)
// Time Complexity: O(n)
func moveZeroes(nums []int) {
	zeroPtr := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroPtr] = nums[i]
			zeroPtr++
		}
	}

	for i := zeroPtr; i < len(nums); i++ {
		nums[i] = 0
	}
}
