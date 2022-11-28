package leetcode

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	noLoseMap := map[int]bool{}
	oneLoseMap := map[int]bool{}
	for i := 0; i < len(matches); i++ {
		winner := matches[i][0]
		if _, exist := noLoseMap[winner]; !exist {
			noLoseMap[winner] = true
		}

		loser := matches[i][1]
		noLoseMap[loser] = false
		if _, exist := oneLoseMap[loser]; !exist {
			oneLoseMap[loser] = true
		} else {
			oneLoseMap[loser] = false
		}
	}

	noLoses := make([]int, 0, len(noLoseMap))
	for key, value := range noLoseMap {
		if value {
			noLoses = append(noLoses, key)
		}
	}
	sort.Ints(noLoses)
	oneLoses := make([]int, 0, len(oneLoseMap))
	for key, value := range oneLoseMap {
		if value {
			oneLoses = append(oneLoses, key)
		}
	}
	sort.Ints(oneLoses)
	return [][]int{noLoses, oneLoses}
}

func Test_findWinners() {
	fmt.Println(findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}})) //[[1 2 10] [4 5 7 8]]
}
