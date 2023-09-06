package leetcode

import "fmt"

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	res := 0
	counter := make(map[int]int, n)
	counter[0] = 1
	preSum := 0
	for _, num := range nums {
		if num%modulo == k {
			preSum++
		}
		remain := (preSum%modulo - k + modulo) % modulo
		res += counter[remain]
		counter[preSum%modulo]++
	}
	return int64(res)
}

func Test_countInterestingSubarrays() {
	fmt.Println(countInterestingSubarrays([]int{3, 2, 4}, 2, 1)) //3
}
