package leetcode

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	queue := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	count := 1
	for len(queue) > 0 {
		length := len(queue)
		for _, cell := range queue[:length] {
			r, c := cell[0], cell[1]
			for _, dir := range dirs {
				newR, newC := r+dir[0], c+dir[1]
				if newR >= 0 && newR < m && newC >= 0 && newC < n && mat[newR][newC] == -1 {
					mat[newR][newC] = count
					queue = append(queue, []int{newR, newC})
				}
			}
		}
		queue = queue[length:]
		count++
	}
	return mat
}

func updateMatrix_dp(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			temp := 10001
			if i > 0 {
				temp = min(temp, mat[i-1][j]+1)
			}
			if j > 0 {
				temp = min(temp, mat[i][j-1]+1)
			}
			mat[i][j] = temp
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}
			if i < m-1 {
				mat[i][j] = min(mat[i][j], mat[i+1][j]+1)
			}
			if j < n-1 {
				mat[i][j] = min(mat[i][j], mat[i][j+1]+1)
			}
		}
	}
	return mat
}
