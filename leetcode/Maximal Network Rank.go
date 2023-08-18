package leetcode

func maximalNetworkRank(n int, roads [][]int) int {
	roadMap := make([][]bool, n)
	for i := range roadMap {
		roadMap[i] = make([]bool, n)
	}
	roadCount := make([]int, n)
	for _, road := range roads {
		i, j := road[0], road[1]
		roadMap[i][j] = true
		roadMap[j][i] = true
		roadCount[i]++
		roadCount[j]++
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if roadMap[i][j] {
				res = max(res, roadCount[i]+roadCount[j]-1)
			} else {
				res = max(res, roadCount[i]+roadCount[j])
			}
		}
	}
	return res
}
