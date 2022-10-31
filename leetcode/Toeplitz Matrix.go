package leetcode

import (
	"fmt"
)

func isToeplitzMatrix(matrix [][]int) bool {
	currRow := matrix[0]
	rowLength := len(currRow)
	for i := 1; i < len(matrix); i++ {
		//can't import slices in leecode
		// if !slices.Equal(currRow[0:rowLength-1], matrix[i][1:rowLength]) {
		// 	return false
		// }

		if !IsSliceEqual(currRow[0:rowLength-1], matrix[i][1:rowLength]) {
			return false
		}
		currRow = matrix[i]
	}
	return true
}

func IsSliceEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_isToeplitzMatrix() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}})) //true
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))                           //false
}
