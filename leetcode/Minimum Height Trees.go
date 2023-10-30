package leetcode

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := make([]map[int]bool, n)
	for i := range graph {
		graph[i] = map[int]bool{}
	}
	for _, e := range edges {
		graph[e[0]][e[1]] = true
		graph[e[1]][e[0]] = true
	}
	leaves := []int{}
	for i, g := range graph {
		if len(g) == 1 {
			leaves = append(leaves, i)
		}
	}
	remain := n
	for remain > 2 {
		remain -= len(leaves)
		nextLeaves := []int{}
		for _, leaf := range leaves {
			for neighbor := range graph[leaf] {
				delete(graph[neighbor], leaf)
				if len(graph[neighbor]) == 1 {
					nextLeaves = append(nextLeaves, neighbor)
				}
			}
		}
		leaves = nextLeaves
	}
	return leaves
}

func Test_findMinHeightTrees() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}})) // [1]
}
