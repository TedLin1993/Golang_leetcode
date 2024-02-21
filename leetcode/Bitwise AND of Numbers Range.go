package leetcode

func rangeBitwiseAnd(left int, right int) int {
	res := right
	for res > left {
		res &= res - 1
	}
	return res
}
