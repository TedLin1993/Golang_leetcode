package leetcode

import (
	"fmt"
)

func minCapability(nums []int, k int) int {
	left, right := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		left = min(left, nums[i])
		right = max(right, nums[i])
	}
	for left < right {
		mid := left + (right-left)/2
		count := 0
		for i := 0; i < len(nums); i++ {
			if mid >= nums[i] {
				count++
				i++
			}
		}
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func Test_minCapability() {
	// fmt.Println(minCapability([]int{2, 3, 5, 9}, 2))
	fmt.Println(minCapability([]int{25038, 3053, 2825, 3638, 4648, 3259, 4948, 4248, 4940, 2834, 109, 5224, 5097, 4708, 2162, 3438, 4152, 4134, 551, 3961, 2294, 3961, 1327, 2395, 1002, 763, 4296, 3147, 5069, 2156, 572, 1261, 4272, 4158, 5186, 2543, 5055, 4735, 2325, 1206, 1019, 1257, 5048, 1563, 3507, 4269, 5328, 173, 5007, 2392, 967, 2768, 86, 3401, 3667, 4406, 4487, 876, 1530, 819, 1320, 883, 1101, 5317, 2305, 89, 788, 1603, 3456, 5221, 1910, 3343, 4597}, 28)) //4134
}
