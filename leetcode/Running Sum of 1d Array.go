package leetcode

import "fmt"

func runningSum(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		nums[i] = sum
	}

	return nums
}

func TestrunningSum() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))     //[1,3,6,10]
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))  //[1,2,3,4,5]
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1})) //[3,4,6,16,17]
}
