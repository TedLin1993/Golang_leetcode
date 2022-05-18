package leetcode

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {
	length := len(grid)
	if grid[0][0] == 1 || grid[length-1][length-1] == 1 {
		return -1
	}
	if length == 1 {
		return 1
	}

	queue := [][]int{{0, 0}}
	layer := 0
	adjacentOffsets := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, -1},
		{-1, 1},
		{-1, -1},
		{1, 1},
	}
	for len(queue) > 0 {
		layer++
		count := len(queue)
		for i := 0; i < count; i++ {
			node := queue[i]
			for _, offset := range adjacentOffsets {
				nextI := node[0] + offset[0]
				nextJ := node[1] + offset[1]
				if nextI == length-1 && nextJ == length-1 {
					return layer + 1
				}

				if isValidCell(nextI, nextJ, length) && grid[nextI][nextJ] == 0 {
					queue = append(queue, []int{nextI, nextJ})
					grid[nextI][nextJ] = 1
				}
			}
		}

		queue = queue[count:]
	}
	return -1
}

func isValidCell(i, j, length int) bool {
	return i >= 0 && i < length && j >= 0 && j < length
}

func TestshortestPathBinaryMatrix() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}))                  //2
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}})) //4
	fmt.Println(shortestPathBinaryMatrix([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}})) //-1
}
