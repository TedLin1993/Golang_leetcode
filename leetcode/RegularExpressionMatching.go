package leetcode

import "fmt"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true

	for j := 1; j < len(p); j++ {
		if p[j] == '*' {
			dp[0][j+1] = dp[0][j-1]
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				if s[i] == p[j-1] || p[j-1] == '.' {
					dp[i+1][j+1] = dp[i+1][j-1] || dp[i+1][j] || dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j-1]
				}
			} else {
				dp[i+1][j+1] = false
			}
		}
	}
	return dp[len(s)][len(p)]
}

func TestRegularExpressionMatching() {
	fmt.Println(isMatch("aa", "a"))                                   //false
	fmt.Println(isMatch("aaa", "ab*a"))                               //false
	fmt.Println(isMatch("abcd", "d*"))                                //false
	fmt.Println(isMatch("aa", "a*"))                                  //true
	fmt.Println(isMatch("ab", ".*"))                                  //true
	fmt.Println(isMatch("abcasdas", ".*"))                            //true
	fmt.Println(isMatch("aab", "c*a*b"))                              //true
	fmt.Println(isMatch("aaa", "a*a"))                                //true
	fmt.Println(isMatch("aaaaaaaa", "a*aa"))                          //true
	fmt.Println(isMatch("aaab", "a*ab"))                              //true
	fmt.Println(isMatch("aa", "a*b*"))                                //true
	fmt.Println(isMatch("aaaa", "aa*a"))                              //true
	fmt.Println(isMatch("sippi", "si.*"))                             //true
	fmt.Println(isMatch("sippi", "s*p*."))                            //false
	fmt.Println(isMatch("ippi", "s*p*."))                             //false
	fmt.Println(isMatch("ba", ".*a*a"))                               //true
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*b*b*")) //true
	fmt.Println(isMatch("aaa", "ab*a*c*a"))                           //true
	fmt.Println(isMatch("aaabaaaababcbccbaab", "c*c*.*c*a*..*c*"))    //true
}
