package leetcode

func getGoodIndices(variables [][]int, target int) []int {
	n := len(variables)
	res := make([]int, 0, n)
	for i, v := range variables {
		a, b, c, m := v[0], v[1], v[2], v[3]
		ab := a
		for i := 1; i < b; i++ {
			ab = ab * a % 10
		}
		ab %= 10
		abc := ab
		for i := 1; i < c; i++ {
			abc = abc * ab % m
		}
		abc %= m
		if abc == target {
			res = append(res, i)
		}
	}
	return res
}
