package game

// Time complexity : O(n)
// Space complexity : O(1)
func canJump(nums []int) bool {
	lastPos := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}
