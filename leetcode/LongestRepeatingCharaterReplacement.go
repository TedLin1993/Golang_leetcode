package leetcode

func characterReplacement(s string, k int) int {
	mapChar := map[int]int{}
	for i := 0; i < 26; i++ {
		char := 'A' + i
		mapChar[char] = 0
	}

	res := 0
	right := 0
	mapChar[int(s[0])]++
	maxCount := 1
	for left := 0; left < len(s); left++ {
		for right-left+1-maxCount <= k {
			if res < right-left+1 {
				res = right - left + 1
			}
			right++
			if right > len(s)-1 {
				break
			}
			mapChar[int(s[right])]++
			maxCount = max(maxCount, mapChar[int(s[right])])

		}
		if right > len(s)-1 {
			break
		}
		mapChar[int(s[left])]--
	}

	return res
}
