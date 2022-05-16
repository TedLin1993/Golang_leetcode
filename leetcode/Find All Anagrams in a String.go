package leetcode

import "fmt"

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	pArray := make([]int, 26) //a-z 26 characters
	for i := 0; i < len(p); i++ {
		pArray[p[i]-'a']++
	}

	currentArray := make([]int, 26)
	//first len(p) char
	for i := 0; i < len(p); i++ {
		currentArray[s[i]-'a']++
	}

	res := []int{}
	for i := len(p); i < len(s); i++ {
		if checkAnagram(pArray, currentArray) {
			res = append(res, i-len(p))
		}
		currentArray[s[i-len(p)]-'a']--
		currentArray[s[i]-'a']++
	}
	if checkAnagram(pArray, currentArray) {
		res = append(res, len(s)-len(p))
	}

	return res
}

func checkAnagram(pArray []int, currentArray []int) bool {
	for i := 0; i < 26; i++ {
		if pArray[i] != currentArray[i] {
			return false
		}
	}
	return true
}

func TestFindAnagrams() {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) //[0,6]
	fmt.Println(findAnagrams("abab", "ab"))        //[0,1,2]
	fmt.Println(findAnagrams("cbaebabac", "abc"))  //[0,6]
}
