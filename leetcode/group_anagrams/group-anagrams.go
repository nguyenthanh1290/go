package group_anagrams

import (
	"sort"
	"strings"
)

// Time Complexity: O(NKlogK)
// Space Complexity: O(NK)
func groupAnagramsSort(strs []string) [][]string {
	dic := make(map[string][]string)

	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		k := strings.Join(s, "")

		dic[k] = append(dic[k], str)
	}

	ang := make([][]string, len(dic))
	i := 0
	for _, v := range dic {
		ang[i] = v
		i++
	}

	return ang
}

// Time Complexity: O(NK)
// Space Complexity: O(NK)
func groupAnagrams(strs []string) [][]string {
	dic := make(map[[26]int][]string)

	for _, str := range strs {
		key := [26]int{}
		for _, r := range []rune(str) {
			idx := r - 97
			key[idx]++
		}

		dic[key] = append(dic[key], str)
	}

	ang := make([][]string, len(dic))
	i := 0
	for _, v := range dic {
		ang[i] = v
		i++
	}

	return ang
}
