package leetcode

import "fmt"

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	trustedCounts := make([]int, n+1)
	isTrustOther := make([]bool, n+1)
	theOne := -1
	for _, t := range trust {
		isTrustOther[t[0]] = true
		trustedCounts[t[1]]++
		if trustedCounts[t[1]] == n-1 {
			theOne = t[1]
		}
	}
	if theOne != -1 && !isTrustOther[theOne] {
		return theOne
	}
	return -1
}

func TestFind_the_Town_Judge() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}
