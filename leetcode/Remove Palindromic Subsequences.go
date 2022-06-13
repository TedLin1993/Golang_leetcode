package leetcode

import "fmt"

func removePalindromeSub(s string) int {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return 2
		}
		left++
		right--
	}
	return 1
}

func TestremovePalindromeSub() {
	fmt.Println(removePalindromeSub("ababa")) //1
	fmt.Println(removePalindromeSub("abb"))   //2
	fmt.Println(removePalindromeSub("baabb")) //2
}
