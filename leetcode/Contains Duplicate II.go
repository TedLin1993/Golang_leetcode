package leetcode

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	digitMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		digit := nums[i]
		index, isExist := digitMap[digit]
		if isExist && i-index <= k {
			return true
		}
		digitMap[digit] = i
	}
	return false
}

func Test_containsNearbyDuplicate() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))       //true
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))       //true
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)) //false
	fmt.Println(containsNearbyDuplicate([]int{-1, -1}, 2))           //true
}
