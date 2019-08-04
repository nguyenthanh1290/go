package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(N int) int {
	str := strconv.FormatInt(int64(N), 2)
	str = strings.TrimRight(str, "0")
	arr := strings.Split(str, "1")

	longest := 0
	for _, v := range arr {
		if len(v) > longest {
			longest = len(v)
		}
	}

	return longest
}

func main() {
	fmt.Println(Solution(2147483646))
}

// https://app.codility.com/programmers/lessons/1-iterations/binary_gap/