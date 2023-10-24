package leetcode

import "fmt"

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	sCount, pCount := [26]int{}, [26]int{}
	for i := 0; i < pLen; i++ {
		pCount[p[i]-'a']++
	}
	res := []int{}
	for i := 0; i < sLen; i++ {
		sCount[s[i]-'a']++
		if i >= pLen {
			sCount[s[i-pLen]-'a']--
		}
		if pCount == sCount {
			res = append(res, i-pLen+1)
		}
	}
	return res
}

func TestFindAnagrams() {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) //[0,6]
	fmt.Println(findAnagrams("abab", "ab"))        //[0,1,2]
	fmt.Println(findAnagrams("cbaebabac", "abc"))  //[0,6]
}
