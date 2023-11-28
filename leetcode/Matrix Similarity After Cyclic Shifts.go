package leetcode

func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	k %= n
	for i := range mat {
		for j, v := range mat[i] {
			idx := j + k
			if idx >= n {
				idx -= n
			}
			if v != mat[i][idx] {
				return false
			}
		}
	}
	return true
}
