package leetcode

func plusOne(digits []int) []int {
	res := digits
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			res[i] = 0
		} else {
			res[i] = digits[i] + 1
			break
		}
		if i == 0 {
			res = append([]int{1}, res...)
		}
	}
	return res
}
