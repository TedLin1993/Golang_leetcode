package leetcode

func halvesAreAlike(s string) bool {
	vowelMap := make([]bool, 'z'-'A'+1)
	vowelMap['a'-'A'] = true
	vowelMap['e'-'A'] = true
	vowelMap['i'-'A'] = true
	vowelMap['o'-'A'] = true
	vowelMap['u'-'A'] = true
	vowelMap['A'-'A'] = true
	vowelMap['E'-'A'] = true
	vowelMap['I'-'A'] = true
	vowelMap['O'-'A'] = true
	vowelMap['U'-'A'] = true

	mid := len(s) / 2
	leftCount := 0
	for i := 0; i < mid; i++ {
		if vowelMap[s[i]-'A'] {
			leftCount++
		}
	}

	rightCount := 0
	for i := mid; i < len(s); i++ {
		if vowelMap[s[i]-'A'] {
			rightCount++
		}
	}

	return leftCount == rightCount
}
