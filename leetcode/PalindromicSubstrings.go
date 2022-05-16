package leetcode

func countSubstrings(s string) int {
	res := 0
	//Palindromic substring of odd length
	for idx := range s {
		left, right := idx, idx
		for left >= 0 && right < len(s) && s[left] == s[right] {
			res++
			left--
			right++
		}
	}
	//Palindromic substring of even length
	for idx := range s {
		left, right := idx, idx+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			res++
			left--
			right++
		}
	}

	return res
}
