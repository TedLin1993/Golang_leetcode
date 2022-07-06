package leetcode

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		}
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return false
	})

	res := [][]int{}
	for _, p := range people {
		idx := p[1]
		if len(res) <= idx {
			res = append(res, p)
		} else {
			res = append(res[:idx+1], res[idx:]...)
			res[idx] = p
		}
	}
	return res
}

func TestreconstructQueue() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})) //[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
	fmt.Println(reconstructQueue([][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}})) //[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
}
