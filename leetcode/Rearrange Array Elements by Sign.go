package leetcode

func rearrangeArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	posIdx, negIdx := 0, 1
	for _, v := range nums {
		if v > 0 {
			res[posIdx] = v
			posIdx += 2
		} else {
			res[negIdx] = v
			negIdx += 2
		}
	}
	return res
}
