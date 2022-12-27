package leetcode

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	stateRest := 0
	stateBuy := -prices[0]
	stateSell := math.MinInt

	for i := 1; i < len(prices); i++ {
		preRest, preBuy, preSell := stateRest, stateBuy, stateSell
		stateRest = max(preRest, preSell)
		stateBuy = max(preBuy, preRest-prices[i])
		stateSell = preBuy + prices[i]
	}
	return max(stateRest, stateSell)
}

func Test_maxProfit() {
	fmt.Println(maxProfit([]int{1}))                      //0
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))          //3
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2, 3, 4, 5})) //6
}
