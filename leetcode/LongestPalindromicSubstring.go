package leetcode

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		lenA := expandCenter(s, i, i)
		lenB := expandCenter(s, i, i+1)
		max := int(math.Max(float64(lenA), float64(lenB)))
		if max > right-left+1 {
			left = i - (max-1)/2
			right = i + max/2
		}
	}
	return s[left : right+1]
}

func expandCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func Test_longestPalindrome() {
	var s string
	s = "babad"
	fmt.Println(longestPalindrome(s)) //bab
	s = "cbbd"
	fmt.Println(longestPalindrome(s)) //bb
	s = "a"
	fmt.Println(longestPalindrome(s)) //a
	s = "ac"
	fmt.Println(longestPalindrome(s)) //a
}
