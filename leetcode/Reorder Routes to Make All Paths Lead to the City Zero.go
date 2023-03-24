package leetcode

import "fmt"

type direction struct {
	to   int
	from int
}

func minReorder(n int, connections [][]int) int {
	connectionMap := make([][]direction, n)
	for _, conn := range connections {
		connectionMap[conn[0]] = append(connectionMap[conn[0]], direction{to: conn[1], from: conn[0]})
		connectionMap[conn[1]] = append(connectionMap[conn[1]], direction{to: conn[0], from: conn[0]})
	}

	queue := connectionMap[0]
	visited := make([]bool, n)
	visited[0] = true
	res := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			dir := queue[i]
			if visited[dir.to] {
				continue
			}
			visited[dir.to] = true
			if dir.to != dir.from {
				res++
			}
			queue = append(queue, connectionMap[dir.to]...)
		}
		queue = queue[length:]
	}
	return res
}

func Test_minReorder() {
	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}})) //3
	fmt.Println(minReorder(5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}))         //2
	fmt.Println(minReorder(3, [][]int{{1, 0}, {2, 0}}))                         //0
}
