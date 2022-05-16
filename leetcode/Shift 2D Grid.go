package leetcode

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	size := m * n

	newGrid := make([][]int, m)
	for i := 0; i < m; i++ {
		newGrid[i] = make([]int, n)
	}

	for i := 0; i < size; i++ {
		index := (i + k) % size
		newGrid[index/n][index%n] = grid[i/n][i%n]
	}

	return newGrid
}

func TestshiftGrid() {
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))    //[[9 1 2] [3 4 5] [6 7 8]]
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9))    //[[1 2 3] [4 5 6] [7 8 9]]
	fmt.Println(shiftGrid([][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}, 23)) //[[6] [5] [1] [2] [3] [4] [7]]

}
