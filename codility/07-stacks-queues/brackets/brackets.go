package brackets

// https://app.codility.com/c/run/trainingNCAFFJ-FF3/
// https://app.codility.com/demo/results/trainingNCAFFJ-FF3/
// Detected time complexity:
// O(N)
func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}

	stack := []rune{}
	for _, v := range S {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')', '}', ']':
			if len(stack) == 0 {
				return 0
			}
			pair := string(stack[len(stack)-1]) + string(v)
			if pair != "()" && pair != "{}" && pair != "[]" {
				return 0
			}
			stack = stack[:len(stack)-1]
		default:
			return 0
		}
	}

	if len(stack) > 0 {
		return 0
	}
	return 1
}
