package leetcode

func characterReplacement(s string, k int) int {
	res := 0
	n := len(s)
	left, right := 0, 0
	maxFreq := 0
	charCount := make([]int, 26)
	for right < n {
		charCount[s[right]-'A']++
		maxFreq = max(maxFreq, charCount[s[right]-'A'])
		for right-left+1-maxFreq > k {
			charCount[s[left]-'A']--
			left++
		}
		res = max(res, right-left+1)
		right++
	}
	return res
}
