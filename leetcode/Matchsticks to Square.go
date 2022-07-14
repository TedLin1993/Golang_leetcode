package leetcode

import (
	"fmt"
	"sort"
)

func makesquare(matchsticks []int) bool {
	sum := 0
	for _, stick := range matchsticks {
		sum += stick
	}
	if sum%4 != 0 {
		return false
	}
	edgeLen := sum / 4
	edge := [4]int{}
	memory := make(map[[4]int]bool)
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	return dfs_makesquare(matchsticks, 0, edgeLen, edge, memory)
}

func dfs_makesquare(matchsticks []int, index int, edgeLen int, edge [4]int, memory map[[4]int]bool) bool {
	if index == len(matchsticks) {
		return edge[0] == edge[1] && edge[1] == edge[2] && edge[2] == edge[3]
	}

	if memory[edge] {
		return false
	}
	memory[edge] = true

	for i := 0; i < 4; i++ {
		if matchsticks[index]+edge[i] <= edgeLen {
			edge[i] += matchsticks[index]
			if dfs_makesquare(matchsticks, index+1, edgeLen, edge, memory) {
				return true
			}

			edge[i] -= matchsticks[index]
			if edge[i] == 0 {
				break
			}
		}
	}
	return false
}

func Testmakesquare() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))                      //true
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))                      //false
	fmt.Println(makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3})) //true
}
