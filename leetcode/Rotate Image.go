package leetcode

import "fmt"

func rotate(matrix [][]int) {
	length := len(matrix)
	layer := 0
	for layer < length/2 {
		for i := layer; i < length-1-layer; i++ {
			lastIndex := length - 1 - layer
			matrix[layer][i], matrix[i][lastIndex], matrix[lastIndex][length-1-i], matrix[length-1-i][layer] =
				matrix[length-1-i][layer], matrix[layer][i], matrix[i][lastIndex], matrix[lastIndex][length-1-i]
		}
		layer++
	}
}

func Testrotate() {
	matrix := [][]int{{}}

	matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix) //[[7 4 1] [8 5 2] [9 6 3]]

	matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Println(matrix) //[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

}
