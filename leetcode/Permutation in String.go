package leetcode

func checkInclusion(s1 string, s2 string) bool {
	charMap := ['z' - 'a' + 1]int{}
	for i := 0; i < len(s1); i++ {
		charMap[s1[i]-'a']++
	}

	left, right := 0, 0
	tempCharMap := ['z' - 'a' + 1]int{}
	for right < len(s2) {
		charIdx := s2[right] - 'a'
		tempCharMap[charIdx]++
		if charMap[charIdx] == 0 {
			tempCharMap = ['z' - 'a' + 1]int{}
			left = right + 1
		} else if tempCharMap[charIdx] > charMap[charIdx] {
			for tempCharMap[charIdx] > charMap[charIdx] {
				delCharIdx := s2[left] - 'a'
				tempCharMap[delCharIdx]--
				left++
			}
		} else if len(s1) == right-left+1 {
			return true
		}
		right++
	}
	return false
}
