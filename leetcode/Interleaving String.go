package leetcode

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n3 != n1+n2 {
		return false
	}
	dp := make([]bool, n2+1)
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			dp[j] = (i == 0 && j == 0) ||
				(i > 0 && dp[j] && s1[i-1] == s3[i+j-1]) ||
				(j > 0 && dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[n2]
}

func TestisInterleave() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac")) //true
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc")) //false
	fmt.Println(isInterleave("", "", ""))                     //true
}
