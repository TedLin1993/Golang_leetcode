package leetcode

import "sort"

func minimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)
	sum := 0
	idx := 0
	res := 0
	for sum < target {
		cur := sum + 1
		for idx < len(coins) && coins[idx] <= cur {
			sum += coins[idx]
			idx++
		}
		if sum < cur {
			sum += cur
			res++
		}
	}
	return res
}
