package leetcode

import "fmt"

func equalPairs(grid [][]int) int {
	n := len(grid)
	rowMap := make(map[[200]int]int, n)
	for _, row := range grid {
		var temp [200]int
		copy(temp[:n], row)
		rowMap[temp]++
	}

	res := 0
	for i := 0; i < n; i++ {
		var temp [200]int
		for j := 0; j < n; j++ {
			temp[j] = grid[j][i]
		}
		if rowMap[temp] > 0 {
			res += rowMap[temp]
		}
	}
	return res
}

func Test_equalPairs() {
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))                        //1
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}})) //3
}
