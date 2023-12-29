package leetcode

import "fmt"

func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	charMap := map[byte]int{}
	for i := 0; i < n; i++ {
		charMap[t[i]]++
	}
	minLen := m + 1
	res := ""
	tCount := 0
	l := 0
	for r := 0; r < m; r++ {
		if _, ok := charMap[s[r]]; !ok {
			continue
		}
		charMap[s[r]]--
		if charMap[s[r]] >= 0 {
			tCount++
		}
		for tCount == n {
			if r-l+1 < minLen {
				minLen = r - l + 1
				res = s[l : r+1]
			}
			if _, ok := charMap[s[l]]; ok {
				charMap[s[l]]++
				if charMap[s[l]] > 0 {
					tCount--
				}
			}
			l++
		}
	}
	return res
}

func TestminWindow() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC")) //BANC
	fmt.Println(minWindow("a", "a"))               //a
	fmt.Println(minWindow("a", "aa"))              //""
}
