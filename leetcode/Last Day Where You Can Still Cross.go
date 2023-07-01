package leetcode

import "fmt"

func latestDayToCross(row int, col int, cells [][]int) int {
	n := len(cells)
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		isCross := bfs_latestDayToCross(row, col, cells, mid)
		if isCross {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func bfs_latestDayToCross(row int, col int, cells [][]int, index int) bool {
	graph := make([][]bool, row)
	for i := range graph {
		graph[i] = make([]bool, col)
	}
	for _, cell := range cells[:index+1] {
		graph[cell[0]-1][cell[1]-1] = true
	}
	queue := [][]int{}
	for i := 0; i < col; i++ {
		if !graph[0][i] {
			queue = append(queue, []int{0, i})
			graph[0][i] = true
		}
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		cell := queue[0]
		for _, dir := range dirs {
			nextR, nextC := cell[0]+dir[0], cell[1]+dir[1]
			if nextR >= 0 && nextR < row && nextC >= 0 && nextC < col && !graph[nextR][nextC] {
				if nextR == row-1 {
					return true
				}
				queue = append(queue, []int{nextR, nextC})
				graph[nextR][nextC] = true
			}
		}
		queue = queue[1:]
	}
	return false
}

func Test_latestDayToCross() {
	fmt.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}))                                                                 //2
	fmt.Println(latestDayToCross(2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}))                                                                 //1
	fmt.Println(latestDayToCross(3, 3, [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}))                         //3
	fmt.Println(latestDayToCross(6, 2, [][]int{{4, 2}, {6, 2}, {2, 1}, {4, 1}, {6, 1}, {3, 1}, {2, 2}, {3, 2}, {1, 1}, {5, 1}, {5, 2}, {1, 2}})) //3
}
