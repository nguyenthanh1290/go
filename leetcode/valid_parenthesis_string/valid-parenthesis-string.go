package string

// Approach: Stack
// Time Complexity: O(N)
// Space Complexity: O(N)
func checkValidString(s string) bool {
	if len(s) == 0 {
		return true
	}

	openStack := []int{}
	starStack := []int{}

	for i, r := range s {
		switch r {
		case '(':
			openStack = append(openStack, i)
		case '*':
			starStack = append(starStack, i)
		case ')':
			if len(openStack) != 0 {
				openStack = openStack[:len(openStack)-1]

			} else if len(starStack) != 0 {
				starStack = starStack[:len(starStack)-1]

			} else {
				return false
			}
		}
	}

	for lgh := len(openStack); lgh != 0; lgh = len(openStack) {
		if len(starStack) == 0 {
			return false

		} else if openStack[lgh-1] < starStack[len(starStack)-1] {
			openStack = openStack[:lgh-1]
			starStack = starStack[:len(starStack)-1]

		} else {
			return false
		}
	}

	return true
}

// Approach: Greedy
// Time Complexity: O(N)
// Space Complexity: O(1)
func checkValidStringGreedy(s string) bool {
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
