package leetcode

func maxVowels(s string, k int) int {
	res, temp := 0, 0
	for i := 0; i < k; i++ {
		if isVowel_maxVowels(s[i]) {
			temp++
			res++
		}
	}
	for right := k; right < len(s); right++ {
		left := right - k
		if isVowel_maxVowels(s[left]) {
			temp--
		}
		if isVowel_maxVowels(s[right]) {
			temp++
			res = max(res, temp)
		}
		left++
	}
	return res
}

func isVowel_maxVowels(a byte) bool {
	if a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' {
		return true
	}
	return false
}
