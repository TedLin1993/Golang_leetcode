package leetcode

import "fmt"

func matrixSumQueries(n int, queries [][]int) int64 {
	res := 0
	visitedRow, visitedCol := make([]bool, n), make([]bool, n)
	remainRow, remainCol := n, n
	for i := len(queries) - 1; i >= 0; i-- {
		query := queries[i]
		typeQ, idx, val := query[0], query[1], query[2]
		if typeQ == 0 && !visitedRow[idx] {
			visitedRow[idx] = true
			res += val * remainCol
			remainRow--
		}
		if typeQ == 1 && !visitedCol[idx] {
			visitedCol[idx] = true
			res += val * remainRow
			remainCol--
		}
	}
	return int64(res)
}

func Test_matrixSumQueries() {
	fmt.Println(matrixSumQueries(3, [][]int{{0, 0, 1}, {1, 2, 2}, {0, 2, 3}, {1, 0, 4}})) //23
}
