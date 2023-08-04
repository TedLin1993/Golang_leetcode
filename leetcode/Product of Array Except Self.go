package leetcode

func productExceptSelf(nums []int) []int {
	product := 1
	zeroCount := 0
	for _, num := range nums {
		if num != 0 {
			product *= num
		} else {
			zeroCount++
		}
	}
	res := make([]int, len(nums))
	if zeroCount > 1 {
		return res
	}
	if zeroCount == 1 {
		for i, num := range nums {
			if num == 0 {
				res[i] = product
				return res
			}
		}
	}
	for i, num := range nums {
		res[i] = product / num
	}

	return res
}
