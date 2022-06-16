package leetcode

import "fmt"

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for idx := range dp {
		dp[idx] = make([]int, len(word2)+1)
	}

	//find max common subsequence
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			idx1, idx2 := i-1, j-1
			if word1[idx1] == word2[idx2] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	maxSubseq := dp[len(word1)][len(word2)]

	//calculate removed char length
	return len(word1) - maxSubseq + len(word2) - maxSubseq
}

func TestminDistance() {
	fmt.Println(minDistance("sea", "eat"))       //2
	fmt.Println(minDistance("leetcode", "etco")) //4
}
