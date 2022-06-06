package leetcode

import "fmt"

func transpose(matrix [][]int) [][]int {
	rowLength := len(matrix)
	columnLength := len(matrix[0])
	res := make([][]int, columnLength)
	for i := 0; i < columnLength; i++ {
		res[i] = make([]int, rowLength)
	}

	for i := 0; i < rowLength; i++ {
		for j := 0; j < columnLength; j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}

func Testtranspose() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) //[[1,4,7],[2,5,8],[3,6,9]]
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))            //[[1,4],[2,5],[3,6]]
}
