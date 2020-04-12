package weight

import (
	"sort"
)

func lastStoneWeight(stones []int) int {
	for total := len(stones); total >= 1; total = len(stones) {
		if total == 1 {
			return stones[0]
		}

		sort.Ints(stones)

		x := stones[total-1]
		y := stones[total-2]
		stones = stones[:total-2]
		if x > y {
			stones = append(stones, x-y)
		}
	}

	return 0
}
