package leetcode

import (
	"fmt"
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	res, currentScore := 0, 0
	left, right := 0, len(tokens)-1
	for left <= right {
		if power < tokens[left] && currentScore == 0 {
			return res
		}

		if power >= tokens[left] {
			power -= tokens[left]
			currentScore++
			left++
			res = max(res, currentScore)
		} else {
			power += tokens[right]
			currentScore--
			right--
		}
	}

	return res
}

func Test_bagOfTokensScore() {
	fmt.Println(bagOfTokensScore([]int{100}, 50))                 //0
	fmt.Println(bagOfTokensScore([]int{100, 200}, 150))           //1
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200)) //2

}
