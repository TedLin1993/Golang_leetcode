package leetcode

import "fmt"

func minCost_(nums []int, x int) int64 {
	inCurredCost := 0
	n := len(nums)
	cost := make([]int, n)
	for i, num := range nums {
		cost[i] = num
	}
	res := sumArr(cost)
	for i := 0; i < n; i++ {
		inCurredCost += x
		nums = append(nums[n-1:n], nums[0:n-1]...)
		for j := 0; j < n; j++ {
			if nums[j] < cost[j] {
				cost[j] = nums[j]
			}
		}
		res = min(res, sumArr(cost)+inCurredCost)
	}
	return int64(res)
}

func Test_minCost_() {
	// fmt.Println(minCost_([]int{20, 1, 15}, 5))                                        //13
	// fmt.Println(minCost_([]int{15, 150, 56, 69, 214, 203}, 42))                       //298
	fmt.Println(minCost_([]int{733, 200, 839, 515, 852, 615, 8, 584, 250, 337}, 537)) //3188
}
