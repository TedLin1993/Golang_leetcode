package leetcode

import (
	"fmt"
)

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		count := 0
		row, col := 0, n-1
		for row < n && col >= 0 {
			if matrix[row][col] > mid {
				col--
			} else {
				count += col + 1
				row++
			}
		}

		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

func TestkthSmallest() {
	fmt.Println(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))                                                                //13
	fmt.Println(kthSmallest([][]int{{-5}}, 1))                                                                                                 //-5
	fmt.Println(kthSmallest([][]int{{1, 2}, {1, 3}}, 3))                                                                                       //2
	fmt.Println(kthSmallest([][]int{{-5, -4}, {-5, -4}}, 2))                                                                                   //-5
	fmt.Println(kthSmallest([][]int{{1, 1, 3, 8, 13}, {4, 4, 4, 8, 18}, {9, 14, 18, 19, 20}, {14, 19, 23, 25, 25}, {18, 21, 26, 28, 29}}, 13)) //18
}
