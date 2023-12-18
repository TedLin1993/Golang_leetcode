package leetcode

import "slices"

func minimumCost(nums []int) int64 {
	n := len(nums)
	slices.Sort(nums)
	isPalindromic := func(v int) bool {
		stack := []int{}
		for v > 0 {
			stack = append(stack, v%10)
			v /= 10
		}
		for i := 0; i < len(stack)/2; i++ {
			if stack[i] != stack[len(stack)-1-i] {
				return false
			}
		}
		return true
	}
	getCost := func(p int) int64 {
		res := 0
		for _, v := range nums {
			res += abs(v - p)
		}
		return int64(res)
	}
	mid := nums[n/2]
	if n%2 == 0 {
		mid = (nums[n/2-1] + nums[n/2]) / 2
	}
	left := mid
	for {
		if isPalindromic(left) {
			break
		}
		left--
	}
	right := mid
	for {
		if isPalindromic(right) {
			break
		}
		right++
	}
	return min(getCost(left), getCost(right))
}
