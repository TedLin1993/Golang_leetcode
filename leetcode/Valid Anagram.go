package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charMap := make(map[rune]int)
	for _, c := range s {
		charMap[c]++
	}

	for _, c := range t {
		charMap[c]--
		if charMap[c] < 0 {
			return false
		}
	}
	return true
}
