package fish

// https://app.codility.com/demo/results/trainingCESQ49-XGY/
// Detected time complexity:
// O(N)
func Solution(A []int, B []int) int {
	survivals := 0
	down := []int{}

	for i, v := range A {
		if B[i] == 0 {
			if len(down) == 0 {
				survivals++
			} else {
				j := len(down) - 1
				for j >= 0 {
					if down[j] > v {
						break
					}
					j--
				}
				if j < 0 {
					down = []int{}
					survivals++
				} else {
					down = down[:j+1]
				}
			}
		} else {
			down = append(down, v)
		}
	}

	survivals += len(down)

	return survivals
}