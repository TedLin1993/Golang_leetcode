package leetcode

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	dicMap := make([][]string, 'z'-'a'+1)
	for _, d := range dictionary {
		firstChar := d[0]
		idx := firstChar - 'a'
		dicMap[idx] = append(dicMap[idx], d)
	}

	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1] + 1
		sIdx := s[i] - 'a'
		for _, d := range dicMap[sIdx] {
			dLen := len(d)
			if i+dLen <= n && s[i:i+dLen] == d {
				dp[i] = min(dp[i], dp[i+len(d)])
			}
		}
	}
	return dp[0]
}

func Test_minExtraChar() {
	fmt.Println(minExtraChar("dwmodizxvvbosxxw", []string{"ox", "lb", "diz", "gu", "v", "ksv", "o", "nuq", "r", "txhe", "e", "wmo", "cehy", "tskz", "ds", "kzbu"})) //7
}
