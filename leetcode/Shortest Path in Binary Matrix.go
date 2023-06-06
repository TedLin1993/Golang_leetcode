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

func shortestPathBinaryMatrix_2(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	queue := [][]int{{0, 0}}
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	step := 1
	for len(queue) > 0 {
		length := len(queue)
		for _, cell := range queue[:length] {
			for _, dir := range dirs {
				r, c := cell[0]+dir[0], cell[1]+dir[1]
				if r == n-1 && c == n-1 {
					return step + 1
				}
				if r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 0 && !visited[r][c] {
					queue = append(queue, []int{r, c})
					visited[r][c] = true
				}
			}
		}
		queue = queue[length:]
		step++
	}
	return -1
}

func TestshortestPathBinaryMatrix() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}))                  //2
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}})) //4
	fmt.Println(shortestPathBinaryMatrix([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}})) //-1
}
