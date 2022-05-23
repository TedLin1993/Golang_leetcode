package leetcode

import "fmt"

func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], c[0])
	}
	res := [][]int{}
	ranks := make([]int, n)
	ranks[0] = 0
	for i := 1; i < n; i++ {
		ranks[i] = -1
	}
	dfs_criticalConnections(graph, -1, 0, 0, ranks, &res)
	return res
}

func dfs_criticalConnections(graph [][]int, preNode int, currentNode int, currentLeyer int, ranks []int, result *[][]int) {
	ranks[currentNode] = currentLeyer
	for i := 0; i < len(graph[currentNode]); i++ {
		nextNode := graph[currentNode][i]
		if nextNode == preNode {
			continue
		}
		if ranks[nextNode] == -1 {
			dfs_criticalConnections(graph, currentNode, nextNode, currentLeyer+1, ranks, result)
			if ranks[nextNode] < ranks[currentNode] {
				ranks[currentNode] = ranks[nextNode]
			}
			if ranks[nextNode] > currentLeyer {
				*result = append(*result, []int{currentNode, nextNode})
			}
		} else if ranks[nextNode] < currentLeyer {
			ranks[currentNode] = ranks[nextNode]
		}
	}
}

func TestcriticalConnections() {
	fmt.Println(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))         //[1,3]
	fmt.Println(criticalConnections(2, [][]int{{0, 1}}))                                 //[0,1]
	fmt.Println(criticalConnections(4, [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 3}, {2, 3}})) //[]
}
