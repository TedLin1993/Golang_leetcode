package leetcode

import "fmt"

func findTheDifference(s string, t string) byte {
	characterCount := 26
	sArray := make([]int, characterCount)
	tArray := make([]int, characterCount)

	for i := 0; i < len(s); i++ {
		sArray[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		tArray[t[i]-'a']++
	}

	for i := 0; i < characterCount; i++ {
		if sArray[i] != tArray[i] {
			return byte(i + 'a')
		}
	}
	return 0
}

func TestFindTheDifference() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}
