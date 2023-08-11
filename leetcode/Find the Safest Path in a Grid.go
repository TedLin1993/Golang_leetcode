package leetcode

import (
	"fmt"
)

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = -1
		}
	}

	queue := make([][]int, 0, n*n)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
				dis[i][j] = 0
			}
		}
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	layer := 0
	for len(queue) > 0 {
		length := len(queue)
		for _, cell := range queue {
			for _, dir := range dirs {
				r, c := cell[0]+dir[0], cell[1]+dir[1]
				if r >= 0 && r < n && c >= 0 && c < n && dis[r][c] == -1 {
					queue = append(queue, []int{r, c})
					dis[r][c] = layer + 1
				}
			}
		}
		layer++
		queue = queue[length:]
	}
	left, right := 0, min(dis[0][0], dis[n-1][n-1])
	for left < right {
		mid := left + (right-left+1)/2
		if isSfOk(dis, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func isSfOk(dis [][]int, sf int) bool {
	n := len(dis)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if r == n-1 && c == n-1 {
			return true
		}
		for _, dir := range dirs {
			nextR, nextC := r+dir[0], c+dir[1]
			if nextR >= 0 && nextR < n && nextC >= 0 && nextC < n && !visited[nextR][nextC] && dis[nextR][nextC] >= sf {
				visited[nextR][nextC] = true
				if dfs(nextR, nextC) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, 0)
}

func Test_maximumSafenessFactor() {
	fmt.Println(maximumSafenessFactor([][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}))                        //2
	fmt.Println(maximumSafenessFactor([][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}})) //2
	fmt.Println(maximumSafenessFactor([][]int{{0, 1, 1}, {0, 1, 1}, {1, 1, 1}}))                        //0
}
