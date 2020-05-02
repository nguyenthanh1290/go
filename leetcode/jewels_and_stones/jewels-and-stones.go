package jewstones

// Time complexity: O(n)
// Space complexity: O(n)
func numJewelsInStones(J string, S string) int {
	jewels := make(map[rune]struct{}, len(S))
	for _, j := range J {
		jewels[j] = struct{}{}
	}

	count := 0
	for _, s := range S {
		if _, ok := jewels[s]; ok {
			count++
		}
	}

	return count
}
