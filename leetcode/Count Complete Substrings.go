package leetcode

import "fmt"

func countCompleteSubstrings(word string, k int) int {
	res := 0
	start := 0
	for start < len(word) {
		end := start
		for i := start + 1; i < len(word); i++ {
			if abs(int(word[i])-int(word[i-1])) <= 2 {
				end = i
			} else {
				break
			}
		}
		res += getComplete(word[start:end+1], k)
		start = end + 1
	}
	return res
}

func getComplete(s string, k int) int {
	res := 0
	for i := 1; i <= 26; i++ {
		size := i * k
		if size > len(s) {
			break
		}
		charCount := make([]int, 26)
		for j := 0; j < size; j++ {
			charCount[s[j]-'a']++
		}
		if isComplete(charCount, k) {
			res++
		}
		for j := size; j < len(s); j++ {
			charCount[s[j]-'a']++
			charCount[s[j-size]-'a']--
			if isComplete(charCount, k) {
				res++
			}
		}
	}
	return res
}

func isComplete(charCount []int, k int) bool {
	res := true
	for _, count := range charCount {
		if count != 0 && count != k {
			res = false
			break
		}
	}
	return res
}

func Test_countCompleteSubstrings() {
	fmt.Println(countCompleteSubstrings("igigee", 2)) //3
}
