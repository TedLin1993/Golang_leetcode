package leetcode

import "fmt"

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	// 1: circle; 2: no circle
	circleMap := make([]int, n)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if circleMap[idx] == 1 {
			return false
		}
		if circleMap[idx] == 2 {
			return true
		}
		if len(graph[idx]) == 0 {
			circleMap[idx] = 2
			return true
		}

		circleMap[idx] = 1
		for _, node := range graph[idx] {
			if !dfs(node) {
				return false
			}
		}
		circleMap[idx] = 2
		return true
	}
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if dfs(i) {
			res = append(res, i)
		}
	}
	return res
}

func Test_eventualSafeNodes() {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))    //[2,4,5,6]
	fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}})) //[4]
	fmt.Println(eventualSafeNodes([][]int{{2, 3}, {2, 3, 4}, {3, 4}, {}, {1}}))       //[3]
}
