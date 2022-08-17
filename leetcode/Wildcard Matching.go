package leetcode

import "fmt"

func isMatch(s string, p string) bool {
	dp := make([]bool, len(p)+1)
	dp[0] = true
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			dp[i+1] = true
		} else {
			break
		}
	}

	for i := 0; i < len(s); i++ {
		nextDp := make([]bool, len(p)+1)
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '?' {
				nextDp[j+1] = dp[j]
			} else if p[j] == '*' {
				nextDp[j+1] = dp[j] || nextDp[j] || dp[j+1]
			} else {
				nextDp[j+1] = false
			}
		}
		dp = nextDp
	}

	return dp[len(p)]
}

func Test_isMatch() {
	fmt.Println(isMatch("aa", "a"))         //false
	fmt.Println(isMatch("aa", "*"))         //true
	fmt.Println(isMatch("cb", "?a"))        //false
	fmt.Println(isMatch("cb", "?b"))        //true
	fmt.Println(isMatch("cbacacac", "?b*")) //true
	fmt.Println(isMatch("adceb", "*a*b"))   //true
	fmt.Println(isMatch("adceb", "**a*b"))  //true
}
