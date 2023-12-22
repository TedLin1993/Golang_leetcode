package leetcode

import "fmt"

func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}

func subarraySum_PreSum(nums []int, k int) int {
	preSumMap := make(map[int]int, len(nums)+1)
	preSumMap[0] = 1
	res := 0
	cur := 0
	for _, v := range nums {
		cur += v
		res += preSumMap[cur-k]
		preSumMap[cur]++
	}
	return res
}

func TestSubarraySum() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
