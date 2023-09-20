package leetcode

func minOperations_1658(nums []int, x int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < x {
		return -1
	}
	if sum == x {
		return n
	}
	res := n
	left := 0
	for i, num := range nums {
		if x < num {
			left = i - 1
			break
		}
		x -= num
	}
	if x == 0 {
		res = left + 1
	}
	for right := n - 1; right > left; right-- {
		x -= nums[right]
		for x < 0 && left >= 0 {
			x += nums[left]
			left--
		}
		if x == 0 {
			res = min(res, left+1+n-right)
		}
		if x < 0 {
			break
		}
	}
	if res == n {
		return -1
	}
	return res
}
