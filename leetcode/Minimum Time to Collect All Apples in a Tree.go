package leetcode

import "fmt"

func minTime(n int, edges [][]int, hasApple []bool) int {
	connections := make([][]int, n)
	for i := 0; i < len(connections); i++ {
		connections[i] = make([]int, 0, 3)
	}
	visited := make([]bool, n)
	visited[0] = true
	for _, edge := range edges {
		connections[edge[0]] = append(connections[edge[0]], edge[1])
		connections[edge[1]] = append(connections[edge[1]], edge[0])
	}

	sum := 0
	if connections[0] != nil {
		for _, v := range connections[0] {
			sum += dfs_minTime(v, connections, hasApple, visited)
		}
	}
	return sum
}

func dfs_minTime(vertex int, connections [][]int, hasApple []bool, visited []bool) int {
	sum := 0
	visited[vertex] = true
	for _, v := range connections[vertex] {
		if !visited[v] {
			sum += dfs_minTime(v, connections, hasApple, visited)
		}
	}

	if hasApple[vertex] || sum > 0 {
		sum += 2
	}
	return sum
}

func Test_minTime() {
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, true, true, false})) //8
	fmt.Println(minTime(4, [][]int{{0, 2}, {0, 3}, {1, 2}}, []bool{false, true, false, false}))                                            //4
}
