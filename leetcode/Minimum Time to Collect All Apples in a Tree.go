package leetcode

import "fmt"

func minTime(n int, edges [][]int, hasApple []bool) int {
	connections := make(map[int][]int, len(edges))
	for _, edge := range edges {
		connections[edge[0]] = append(connections[edge[0]], edge[1])
	}

	sum := 0
	if connections[0] != nil {
		for _, v := range connections[0] {
			sum += dfs_minTime(v, connections, hasApple)
		}
	}
	return sum
}

func dfs_minTime(vertex int, connections map[int][]int, hasApple []bool) int {
	sum := 0
	if connections[vertex] != nil {
		for _, v := range connections[vertex] {
			sum += dfs_minTime(v, connections, hasApple)
		}
	}

	if hasApple[vertex] || sum > 0 {
		return sum + 2
	}
	return 0
}

func Test_minTime() {
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, true, true, false})) //8
}
