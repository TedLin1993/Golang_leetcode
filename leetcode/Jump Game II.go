package leetcode

func jump(nums []int) int {
	count := 0
	preMax, curMax := 0, 0
	for curMax < len(nums)-1 {
		next := curMax
		for i := preMax; i <= curMax; i++ {
			if i+nums[i] > next {
				next = i + nums[i]
			}
		}
		preMax = curMax + 1
		curMax = next
		count++
	}
	return count
}
