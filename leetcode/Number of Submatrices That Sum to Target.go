package leetcode

import "fmt"

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	for i := 0; i < len(matrix); i++ {
		matrix[i] = append([]int{0}, matrix[i]...)
	}

	res := 0
	for left := 0; left < len(matrix[0])-1; left++ {
		for right := left + 1; right < len(matrix[0]); right++ {
			area := 0
			accumulate := map[int]int{0: 1}
			for row := 0; row < len(matrix); row++ {
				area += matrix[row][right] - matrix[row][left]
				res += accumulate[area-target]
				accumulate[area] += 1
			}
		}
	}
	return res
}

func TestnumSubmatrixSumTarget() {
	fmt.Println(numSubmatrixSumTarget([][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0)) //4
	fmt.Println(numSubmatrixSumTarget([][]int{{1, -1}, {-1, 1}}, 0))                //5
	fmt.Println(numSubmatrixSumTarget([][]int{{904}}, 0))                           //0
}
