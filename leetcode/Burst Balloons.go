package leetcode

import "fmt"

//猜測: 去掉頭尾後選最小值，直到只剩頭尾為止
// -> fail，[1,2,3,4,5] 最佳解是3->4->2->1->5
var coinSum [][]int

func maxCoins(nums []int) int {
	numbers := []int{}
	numbers = append(numbers, 1)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			numbers = append(numbers, nums[i])
		}
	}
	numbers = append(numbers, 1)

	coinSum = make([][]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		coinSum[i] = make([]int, len(numbers))
	}
	return dcMaxCoins(numbers, 0, len(numbers)-1)
}

func dcMaxCoins(nums []int, left, right int) int {
	if left+1 >= right {
		return 0
	}
	if coinSum[left][right] > 0 {
		return coinSum[left][right]
	}
	max := 0
	for i := left + 1; i < right; i++ {
		temp := nums[left]*nums[i]*nums[right] + dcMaxCoins(nums, left, i) + dcMaxCoins(nums, i, right)
		if temp > max {
			max = temp
		}
	}
	coinSum[left][right] = max
	return max
}

func TestBurst_Balloons() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
	fmt.Println(maxCoins([]int{1, 5}))
}
