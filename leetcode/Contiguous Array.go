package leetcode

func findMaxLength(nums []int) int {
	numMap := map[int]int{0: -1}
	res := 0
	cur := 0
	for i, num := range nums {
		if num == 0 {
			cur--
		} else {
			cur++
		}
		if prei, ok := numMap[cur]; ok {
			res = max(res, i-prei)
		} else {
			numMap[cur] = i
		}
	}
	return res
}
