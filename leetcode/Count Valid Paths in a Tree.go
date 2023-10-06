package leetcode

import "fmt"

func countPaths(n int, edges [][]int) int64 {
	graph := make([][]int, n+1)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	isPrimeList := getPrimeList(n)

	var dfs func(node, prev int) []int
	dfs = func(node, prev int) []int {
		res := []int{node}
		for _, next := range graph[node] {
			if !isPrimeList[next] && next != prev {
				nodes := dfs(next, node)
				res = append(res, nodes...)
			}
		}
		return res
	}

	size := make([]int, n+1)
	res := 0
	for i := 1; i <= n; i++ {
		if !isPrimeList[i] {
			continue
		}
		sum := 0
		for _, j := range graph[i] {
			if isPrimeList[j] {
				continue
			}
			if size[j] == 0 {
				nodes := dfs(j, -1)
				for _, node := range nodes {
					size[node] = len(nodes)
				}
			}
			res += size[j] * sum
			sum += size[j]
		}
		res += sum
	}
	return int64(res)
}

func getPrimeList(n int) []bool {
	list := make([]bool, n+1)
	for i := 2; i < len(list); i++ {
		list[i] = true
	}
	for i := 2; i < len(list); i++ {
		for j := i * i; j < len(list); j += i {
			list[j] = false
		}
	}
	return list
}

func Test_countPaths() {
	// fmt.Println(countPaths(5, [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}}))         //4
	fmt.Println(countPaths(6, [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 5}, {3, 6}})) //6
}
