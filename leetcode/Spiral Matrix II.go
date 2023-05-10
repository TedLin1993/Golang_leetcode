package leetcode

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIdx := 0
	r, c := 0, 0
	current := 1
	for current <= n*n {
		res[r][c] = current
		next_r, next_c := r+directions[dirIdx][0], c+directions[dirIdx][1]
		if next_r < 0 || next_r >= n || next_c < 0 || next_c >= n || res[next_r][next_c] != 0 {
			dirIdx++
			if dirIdx == 4 {
				dirIdx = 0
			}
			next_r, next_c = r+directions[dirIdx][0], c+directions[dirIdx][1]
		}
		r, c = next_r, next_c
		current++
	}
	return res
}
