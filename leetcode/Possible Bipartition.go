package leetcode

import "fmt"

var groupPB []int

func possibleBipartition(n int, dislikes [][]int) bool {
	dislikeMap := make([][]int, n+1)
	for i := 0; i < len(dislikes); i++ {
		node1, node2 := dislikes[i][0], dislikes[i][1]
		dislikeMap[node1] = append(dislikeMap[node1], node2)
		dislikeMap[node2] = append(dislikeMap[node2], node1)
	}

	groupPB = make([]int, n+1)
	for node := 0; node <= n; node++ {
		for _, neighbor := range dislikeMap[node] {
			if findPB(node) == findPB(neighbor) {
				return false
			}
			unionPB(dislikeMap[node][0], neighbor)
		}
	}
	return true
}

func findPB(node int) int {
	if groupPB[node] == 0 {
		groupPB[node] = node
	}

	if groupPB[node] != node {
		groupPB[node] = findPB(groupPB[node])
	}
	return groupPB[node]
}

func unionPB(node1 int, node2 int) {
	node1 = findPB(node1)
	node2 = findPB(node2)
	if node1 == node2 {
		return
	}
	if node2 < node1 {
		groupPB[node1] = node2
	} else {
		groupPB[node2] = node1
	}
}

func Test_possibleBipartition() {
	fmt.Println(possibleBipartition(4, [][]int{{1, 4}, {1, 3}, {2, 4}}))                 // true
	fmt.Println(possibleBipartition(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}})) // false
	fmt.Println(possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))                 // false
}
