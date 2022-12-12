package leetcode

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	word1Map := make([]int, 26)
	word2Map := make([]int, 26)
	for _, w := range word1 {
		word1Map[w-'a']++
	}
	for _, w := range word2 {
		if word1Map[w-'a'] == 0 {
			return false
		}
		word2Map[w-'a']++
	}

	sort.Ints(word1Map)
	sort.Ints(word2Map)
	for i := range word1Map {
		if word1Map[i] != word2Map[i] {
			return false
		}
	}
	return true
}

func Test_closeStrings() {
	fmt.Println(closeStrings("abc", "bca"))       //true
	fmt.Println(closeStrings("a", "aa"))          //false
	fmt.Println(closeStrings("cabbba", "abbccc")) //true
}
