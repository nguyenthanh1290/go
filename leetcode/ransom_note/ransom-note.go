package ransom

// Time complexity: O(n+m)
// Space complexity: O(m)
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	dict := make(map[rune]int)
	for _, r := range magazine {
		dict[r]++
	}

	for _, r := range ransomNote {
		if count, ok := dict[r]; !ok || count <= 0 {
			return false

		} else {
			dict[r]--
		}
	}

	return true
}
