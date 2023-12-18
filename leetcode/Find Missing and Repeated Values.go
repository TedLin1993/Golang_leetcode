package leetcode

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	visited := make([]bool, n*n+1)
	a := -1
	for _, r := range grid {
		for _, v := range r {
			if visited[v] {
				a = v
			} else {
				visited[v] = true
			}
		}
	}
	b := -1
	for i := 1; i <= n*n; i++ {
		if !visited[i] {
			b = i
			break
		}
	}
	return []int{a, b}
}
