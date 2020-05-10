package string

// Time complexity : O(N)
// Space complexity : O(1)
func firstUniqChar(s string) int {
	count := make(map[rune]int)

	for _, r := range s {
		count[r]++
	}

	for i, r := range s {
		if c := count[r]; c == 1 {
			return i
		}
	}

	return -1
}
