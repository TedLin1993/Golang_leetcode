package leetcode

import (
	"fmt"
	"math"
)

func minimumMoves(grid [][]int) int {
	from := make([][]int, 0, 9)
	to := make([][]int, 0, 9)
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 {
				to = append(to, []int{r, c})
			}
			if grid[r][c] > 1 {
				for i := 0; i < grid[r][c]-1; i++ {
					from = append(from, []int{r, c})
				}
			}
		}
	}

	res := math.MaxInt
	n := len(from)
	permutation := getPermutation(n)
	for _, p := range permutation {
		temp := 0
		for tIdx, fIdx := range p {
			temp += abs(from[fIdx][0]-to[tIdx][0]) + abs(from[fIdx][1]-to[tIdx][1])
		}
		res = min(res, temp)
	}
	return res
}

func getPermutation(n int) [][]int {
	res := [][]int{}
	visit := make([]bool, n)
	var dfs func(cur []int)
	dfs = func(cur []int) {
		if len(cur) == n {
			temp := make([]int, n)
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if !visit[i] {
				visit[i] = true
				dfs(append(cur, i))
				visit[i] = false
			}
		}
	}
	dfs([]int{})
	return res
}

func Test_minimumMoves() {
	fmt.Println(minimumMoves([][]int{{7, 0, 0}, {0, 0, 0}, {0, 2, 0}}))
}
