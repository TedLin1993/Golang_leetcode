package leetcode

import (
	"fmt"
	"math"
)

func maxSubarraySumCircular(nums []int) int {
	total := 0
	currentMaxSum := 0
	maxSum := math.MinInt
	currentMinSum := 0
	minSum := math.MaxInt
	for i := 0; i < len(nums); i++ {
		total += nums[i]

		currentMaxSum = max(currentMaxSum+nums[i], nums[i])
		maxSum = max(maxSum, currentMaxSum)

		currentMinSum = min(currentMinSum+nums[i], nums[i])
		minSum = min(minSum, currentMinSum)
	}

	if minSum < total {
		return max(maxSum, total-minSum)
	}
	//every element in nums is less than 0
	return maxSum
}

func Test_maxSubarraySumCircular() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))                        //3
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))                            //10
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))                          //-2
	fmt.Println(maxSubarraySumCircular([]int{1, 2, 3}))                             //6
	fmt.Println(maxSubarraySumCircular([]int{-1, 3, -3, 9, -6, 8, -5, -5, -6, 10})) //20
}
