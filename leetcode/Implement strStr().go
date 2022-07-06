package leetcode

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			break
		}
		isMatch := true
		if haystack[i] == needle[0] {
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					isMatch = false
					break
				}
			}
			if isMatch {
				return i
			}
		}
	}
	return -1
}

func TeststrStr() {
	fmt.Println(strStr("a", "a"))       //0
	fmt.Println(strStr("hello", "ll"))  //2
	fmt.Println(strStr("aaaaa", "bba")) //-1
}
