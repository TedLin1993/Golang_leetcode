package leetcode

func firstPalindrome(words []string) string {
	for _, w := range words {
		isPalindrome := true
		n := len(w)
		for i := 0; i < n/2; i++ {
			if w[i] != w[n-1-i] {
				isPalindrome = false
				break
			}
		}
		if isPalindrome {
			return w
		}
	}
	return ""
}
