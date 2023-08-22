package leetcode

func isAcronym(words []string, s string) bool {
	m, n := len(words), len(s)
	if m != n {
		return false
	}

	for i := 0; i < m; i++ {
		if words[i][0] != s[i] {
			return false
		}
	}
	return true
}
