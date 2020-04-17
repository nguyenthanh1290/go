package string

// Approach #3: Greedy
// Time Complexity: O(N)
// Space Complexity: O(1)
func checkValidString(s string) bool {
	lo := 0
	hi := 0

	for _, r := range s {
		if r == '(' {
			lo++

		} else {
			lo--
		}

		if r != ')' {
			hi++

		} else {
			hi--
		}

		if hi < 0 {
			break
		}

		if lo < 0 {
			lo = 0
		}
	}

	return lo == 0
}
