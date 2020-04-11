package strcompare

// Approach #1: Build String
// Time Complexity: O(M + N), where M, NM,N are the lengths of S and T respectively.
// Space Complexity: O(M + N).
func backspaceCompareStack(S string, T string) bool {
	stackS := buildStack(S)
	stackT := buildStack(T)

	if len(stackS) != len(stackT) {
		return false

	} else {
		for i := 0; i < len(stackS); i++ {
			if stackS[i] != stackT[i] {
				return false
			}
		}
	}

	return true
}

func buildStack(str string) []rune {
	stack := make([]rune, 0)

	for _, r := range str {
		if r == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

		} else {
			stack = append(stack, r)
		}
	}

	return stack
}

// Approach #2: Two Pointer
// Time Complexity: O(M + N), where M, NM,N are the lengths of S and T respectively.
// Space Complexity: O(1).
