package leetcode

import "fmt"

func findSmallestInteger(nums []int, value int) int {
	valMap := make([]int, value)
	for i := 0; i < len(nums); i++ {
		nums[i] %= value
		if nums[i] < 0 {
			nums[i] += value
		}
		valMap[nums[i]]++
	}

	stop := 0
	for i := 0; i < len(valMap); i++ {
		if valMap[i] < valMap[stop] {
			stop = i
		}
	}
	return value*valMap[stop] + stop
}

func Test_findSmallestInteger() {
	fmt.Println(findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 5))                              //4
	fmt.Println(findSmallestInteger([]int{1, -10, 7, 13, 6, 8}, 7))                              //2
	fmt.Println(findSmallestInteger([]int{-11, 0, 1, 2, -17, -16, -15, -14, 3323, 321, 424}, 7)) //8

}
