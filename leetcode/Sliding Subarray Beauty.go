package leetcode

func getSubarrayBeauty(nums []int, k int, x int) []int {
	numMap := make(map[int]int, 101)
	for i := 0; i < k; i++ {
		numMap[nums[i]]++
	}
	n := len(nums)
	res := make([]int, n-k+1)
	res[0] = getBeauty(numMap, x)
	for i := 1; i < n-k+1; i++ {
		numMap[nums[i-1]]--
		numMap[nums[i+k-1]]++
		res[i] = getBeauty(numMap, x)
	}
	return res
}

func getBeauty(numMap map[int]int, x int) int {
	for i := -50; i <= 0; i++ {
		if numMap[i] >= x {
			return i
		}
		x -= numMap[i]
	}
	return 0
}
