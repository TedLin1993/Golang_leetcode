package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, right := 0, 0
	charMap := map[byte]int{}
	res := 0
	for right <= len(s)-1 {
		if _, ok := charMap[s[right]]; ok && charMap[s[right]] >= left {
			left = charMap[s[right]] + 1
		} else if right-left+1 > res {
			res = right - left + 1
		}
		charMap[s[right]] = right
		right++
	}

	return res
}
