package leetcode

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}

	index := 0
	for i := 0; index < len(s) && i < len(t); i++ {
		if t[i] == s[index] {
			index++
		}
	}
	return index == len(s)
}

func TestisSubsequence() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) //true
	fmt.Println(isSubsequence("axc", "ahbgdc")) //false
	fmt.Println(isSubsequence("", "abc"))       //true
	fmt.Println(isSubsequence("b", "abc"))      //true
}
