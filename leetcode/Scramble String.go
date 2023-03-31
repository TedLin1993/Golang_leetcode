package leetcode

import "fmt"

var scrambleMap map[string]bool

func isScramble(s1 string, s2 string) bool {
	scrambleMap = make(map[string]bool)
	return dfs_isScramble(s1, s2)
}

func dfs_isScramble(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if countChar(s1) != countChar(s2) {
		return false
	}

	key := s1 + "_" + s2
	if val, ok := scrambleMap[key]; ok {
		return val
	}
	len := len(s1)
	for i := 1; i < len; i++ {
		if (dfs_isScramble(s1[:i], s2[:i]) && dfs_isScramble(s1[i:], s2[i:])) ||
			(dfs_isScramble(s1[i:], s2[:len-i]) && dfs_isScramble(s1[:i], s2[len-i:])) {
			scrambleMap[key] = true
			return true
		}
	}
	scrambleMap[key] = false
	return false
}

func countChar(s string) [26]int {
	var res [26]int
	for i := 0; i < len(s); i++ {
		res[s[i]-'a']++
	}
	return res
}

func Test_isScramble() {
	// fmt.Println(isScramble("great", "rgeat"))                                             //true
	// fmt.Println(isScramble("abcde", "caebd"))                                             //false
	// fmt.Println(isScramble("a", "a"))                                                     //true
	fmt.Println(isScramble("eebaacbcbcadaaedceaaacadccd", "eadcaacabaddaceacbceaabeccd")) //false
}
