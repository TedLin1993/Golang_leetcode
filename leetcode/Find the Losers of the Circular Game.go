package leetcode

func circularGameLosers(n int, k int) []int {
	visited := make([]bool, n)
	index := 0
	iteration := 1
	for !visited[index] {
		visited[index] = true
		index = (index + iteration*k) % n
		iteration++
	}
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			res = append(res, i+1)
		}
	}
	return res
}
