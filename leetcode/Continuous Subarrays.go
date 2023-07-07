package leetcode

import "fmt"

func continuousSubarrays(nums []int) int64 {
	left := 0
	minIdx, maxIdx := 0, 0
	res := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] <= nums[minIdx] {
			minIdx = right
		}
		if nums[right] >= nums[maxIdx] {
			maxIdx = right
		}
		for nums[maxIdx]-nums[minIdx] > 2 {
			if right == minIdx {
				left = maxIdx + 1
				maxIdx++
				for i := maxIdx + 1; i <= right; i++ {
					if nums[i] >= nums[maxIdx] {
						maxIdx = i
					}
				}
			} else {
				left = minIdx + 1
				minIdx++
				for i := minIdx + 1; i <= right; i++ {
					if nums[i] <= nums[minIdx] {
						minIdx = i
					}
				}
			}
		}

		res += right - left + 1
	}
	return int64(res)
}

func Test_continuousSubarrays() {
	fmt.Println(continuousSubarrays([]int{5, 4, 2, 4})) //8
	fmt.Println(continuousSubarrays([]int{7, 6, 5, 1})) //7
}
