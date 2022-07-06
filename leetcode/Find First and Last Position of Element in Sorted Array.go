package leetcode

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	targetleftIdx := findTargetLeftIdx(nums, target, 0, len(nums)-1)
	if targetleftIdx == -1 {
		return []int{-1, -1}
	}
	targetRightIdx := findTargetRightIdx(nums, target, targetleftIdx, len(nums)-1)

	return []int{targetleftIdx, targetRightIdx}
}

func findTargetLeftIdx(nums []int, target, left, right int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func findTargetRightIdx(nums []int, target, left, right int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && nums[right] == target {
		return right
	}
	return -1
}

func TestsearchRange() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) //[3 4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) //[-1 -1]
	fmt.Println(searchRange([]int{}, 0))                  //[-1 -1]
	fmt.Println(searchRange([]int{2, 2}, 3))              //[-1 -1]
}
