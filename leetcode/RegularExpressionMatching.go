package leetcode

import "fmt"

func isMatch_REM(s string, p string) bool {
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
	fmt.Println(isMatch_REM("aa", "a"))                                   //false
	fmt.Println(isMatch_REM("aaa", "ab*a"))                               //false
	fmt.Println(isMatch_REM("abcd", "d*"))                                //false
	fmt.Println(isMatch_REM("aa", "a*"))                                  //true
	fmt.Println(isMatch_REM("ab", ".*"))                                  //true
	fmt.Println(isMatch_REM("abcasdas", ".*"))                            //true
	fmt.Println(isMatch_REM("aab", "c*a*b"))                              //true
	fmt.Println(isMatch_REM("aaa", "a*a"))                                //true
	fmt.Println(isMatch_REM("aaaaaaaa", "a*aa"))                          //true
	fmt.Println(isMatch_REM("aaab", "a*ab"))                              //true
	fmt.Println(isMatch_REM("aa", "a*b*"))                                //true
	fmt.Println(isMatch_REM("aaaa", "aa*a"))                              //true
	fmt.Println(isMatch_REM("sippi", "si.*"))                             //true
	fmt.Println(isMatch_REM("sippi", "s*p*."))                            //false
	fmt.Println(isMatch_REM("ippi", "s*p*."))                             //false
	fmt.Println(isMatch_REM("ba", ".*a*a"))                               //true
	fmt.Println(isMatch_REM("abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*b*b*")) //true
	fmt.Println(isMatch_REM("aaa", "ab*a*c*a"))                           //true
	fmt.Println(isMatch_REM("aaabaaaababcbccbaab", "c*c*.*c*a*..*c*"))    //true
}
