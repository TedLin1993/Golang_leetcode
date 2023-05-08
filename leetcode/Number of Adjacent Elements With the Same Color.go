package leetcode

func colorTheArray(n int, queries [][]int) []int {
	nums := make([]int, n)
	res := make([]int, len(queries))
	current := 0
	for idx, query := range queries {
		numIdx, color := query[0], query[1]
		if nums[numIdx] != 0 && numIdx > 0 && nums[numIdx] == nums[numIdx-1] {
			current--
		}
		if nums[numIdx] != 0 && numIdx < n-1 && nums[numIdx] == nums[numIdx+1] {
			current--
		}

		nums[numIdx] = color
		if numIdx > 0 && nums[numIdx] == nums[numIdx-1] {
			current++
		}
		if numIdx < n-1 && nums[numIdx] == nums[numIdx+1] {
			current++
		}
		res[idx] = current
	}
	return res
}
