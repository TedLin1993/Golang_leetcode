package leetcode

import "fmt"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	wordIdx1, wordIdx2 := 0, 0
	charIdx1, charIdx2 := 0, 0
	for wordIdx1 < len(word1) && wordIdx2 < len(word2) {
		if charIdx1 == len(word1[wordIdx1]) && charIdx2 == len(word2[wordIdx2]) {
			charIdx1 = 0
			charIdx2 = 0
			wordIdx1++
			wordIdx2++
		} else if charIdx1 == len(word1[wordIdx1]) {
			charIdx1 = 0
			wordIdx1++
		} else if charIdx2 == len(word2[wordIdx2]) {
			charIdx2 = 0
			wordIdx2++
		} else if word1[wordIdx1][charIdx1] == word2[wordIdx2][charIdx2] {
			charIdx1++
			charIdx2++
		} else {
			return false
		}
	}
	return wordIdx1 == len(word1) && wordIdx2 == len(word2)
}

func Test_arrayStringsAreEqual() {
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))          //true
	fmt.Println(arrayStringsAreEqual([]string{"ac", "b"}, []string{"a", "bc"}))          //false
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddef"})) //false
}
