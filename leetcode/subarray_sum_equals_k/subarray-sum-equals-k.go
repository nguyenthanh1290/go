package subarray

// Approach: prefix sum
// Time complexity : O(n^2)
// Space complexity : O(n)
func subarraySumSpaceOn(nums []int, k int) int {
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	count := 0

	for start := 0; start < len(prefixSum)-1; start++ {
		for end := start + 1; end < len(prefixSum); end++ {
			if prefixSum[end]-prefixSum[start] == k {
				count++
			}
		}
	}

	return count
}

// Approach: prefix sum
// Time complexity : O(n^2)
// Space complexity : O(1)
func subarraySumSpaceO1(nums []int, k int) int {
	count := 0

	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}

	return count
}

// Approach: Hashmap
// Time complexity : O(n)
// Space complexity : O(n)
func subarraySum(nums []int, k int) int {
	result := 0
	sum := 0
	f := make(map[int]int)

	f[0] = 1

	for _, num := range nums {
		sum += num
		result += f[sum-k]
		f[sum]++
	}

	return result
}
