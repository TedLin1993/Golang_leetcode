package leetcode

import "fmt"

func spiralOrder(matrix [][]int) []int {
	rowCount, colCount := len(matrix), len(matrix[0])
	res := make([]int, 0, rowCount*colCount)
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	r, c := 0, 0
	dirIdx := 0
	res = append(res, matrix[0][0])
	matrix[0][0] = 101
	for len(res) < rowCount*colCount {
		next_r, next_c := r+dir[dirIdx][0], c+dir[dirIdx][1]
		if next_r >= rowCount || next_r < 0 || next_c >= colCount || next_c < 0 || matrix[next_r][next_c] == 101 {
			dirIdx += 1
			if dirIdx == 4 {
				dirIdx = 0
			}
			next_r, next_c = r+dir[dirIdx][0], c+dir[dirIdx][1]
		}
		r, c = next_r, next_c
		res = append(res, matrix[r][c])
		matrix[r][c] = 101 //exceed the limit
	}
	return res
}

func Test_spiralOrder() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))             //[1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})) //[1,2,3,4,8,12,11,10,9,5,6,7]
}
