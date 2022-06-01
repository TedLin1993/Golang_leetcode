package leetcode

import "fmt"

var charResultMap []int
var charMap []int

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	charResultMap = make([]int, 128)
	charMap = make([]int, 128)
	for i := 0; i < len(t); i++ {
		charResultMap[t[i]]++
	}

	//get first substring reaching result
	leftIndex, rightIndex := 0, 0
	resultCount := 0
	for rightIndex < len(s) {
		if charResultMap[s[rightIndex]] > 0 {
			charMap[s[rightIndex]]++
			if charMap[s[rightIndex]] <= charResultMap[s[rightIndex]] {
				resultCount++
			}
			if resultCount == len(t) {
				break
			}
		}
		rightIndex++
	}
	if resultCount < len(t) {
		return ""
	}
	leftIndex = getLeftIndex(s, leftIndex, rightIndex)

	//sliding windows to get result
	res := s[leftIndex : rightIndex+1]
	rightIndex++
	for rightIndex < len(s) {
		if charResultMap[s[rightIndex]] > 0 {
			charMap[s[rightIndex]]++
			if s[rightIndex] == s[leftIndex] {
				leftIndex = getLeftIndex(s, leftIndex, rightIndex)
				if rightIndex-leftIndex+1 < len(res) {
					res = s[leftIndex : rightIndex+1]
				}
			}
		}
		rightIndex++
	}

	return res
}

func getLeftIndex(s string, leftIndex int, rightIndex int) int {
	for leftIndex <= rightIndex {
		if charResultMap[s[leftIndex]] > 0 {
			if charMap[s[leftIndex]] != charResultMap[s[leftIndex]] {
				charMap[s[leftIndex]]--
			} else {
				break
			}
		}
		leftIndex++
	}
	return leftIndex
}

func TestminWindow() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC")) //BANC
	fmt.Println(minWindow("a", "a"))               //a
	fmt.Println(minWindow("a", "aa"))              //""
}
