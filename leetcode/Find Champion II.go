package leetcode

func findChampion_II(n int, edges [][]int) int {
	isWeak := make([]bool, n)
	for _, e := range edges {
		isWeak[e[1]] = true
	}
	res := -1
	for i, v := range isWeak {
		if !v {
			if res != -1 {
				return -1
			}
			res = i
		}
	}
	return res
}
