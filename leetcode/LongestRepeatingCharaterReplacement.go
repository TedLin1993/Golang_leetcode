package leetcode

func characterReplacement(s string, k int) int {
	mapChar := map[int]int{}
	for i := 0; i < 26; i++ {
		char := 'A' + i
		mapChar[char] = 0
	}

	res := 0
	r := 0
	mapChar[int(s[0])]++
	maxCount := 1
	for l := 0; l < len(s); l++ {
		for r+1-l-maxCount <= k {
			if res < r+1-l {
				res = r + 1 - l
			}
			r++
			if r > len(s)-1 {
				break
			}
			mapChar[int(s[r])]++
			maxCount = max(maxCount, mapChar[int(s[r])])

		}
		if r > len(s)-1 {
			break
		}
		mapChar[int(s[l])]--
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
