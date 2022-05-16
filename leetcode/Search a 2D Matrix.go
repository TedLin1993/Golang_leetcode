package leetcode

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rowLen := len(matrix)
	colLen := len(matrix[0])
	low, high := 0, rowLen*colLen-1
	for low <= high {
		mid := low + (high-low)/2
		if target == matrix[mid/colLen][mid%colLen] {
			return true
		} else if target < matrix[mid/colLen][mid%colLen] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func TestsearchMatrix() {
	// fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 1))  // true
	// fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 11)) // true
	// fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 23)) // true
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 30)) // true
	// fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 31)) // false
	// fmt.Println(searchMatrix([][]int{{1}}, 2))                                               // false

}
