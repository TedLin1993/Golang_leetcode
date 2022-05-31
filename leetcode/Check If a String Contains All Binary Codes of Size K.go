package leetcode

import "fmt"

func hasAllCodes(s string, k int) bool {
	codeMap := make(map[string]bool, k)
	needCount := 1 << k
	for len(s) >= k {
		codeMap[s[:k]] = true
		s = s[1:]
		if len(codeMap) == needCount {
			return true
		}
	}
	return false
}

func TesthasAllCodes() {
	fmt.Println(hasAllCodes("00110110", 2)) // true
	// fmt.Println(hasAllCodes("0110", 1))     // true
}
