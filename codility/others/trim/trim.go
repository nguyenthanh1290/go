package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solution("The quick brown fox jumps over the lazy dog", 39))
}

func Solution(message string, K int) string {
	if len(message) < K {
		return message
	}

	arr := strings.Split(message, " ")
	for i := len(arr); i > 1; i-- {
		result := strings.Join(arr[:i], " ")
		if len(result) <= K {
			return result
		}
	}

	return ""
}
