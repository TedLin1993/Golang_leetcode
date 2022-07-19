package leetcode

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		for j := 1; j < i; j++ {
			temp[j] = res[i-1][j-1] + res[i-1][j]
		}
		temp[i] = 1
		res = append(res, temp)
	}

	return res
}

func Testgenerate() {
	fmt.Println(generate(5))
}
