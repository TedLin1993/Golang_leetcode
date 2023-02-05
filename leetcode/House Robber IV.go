package leetcode

import "fmt"

var maxRobList []int

func minCapability(nums []int, k int) int {
	maxRobList = []int{}
	for i := 0; i < len(nums); i++ {
		dfsGetMaxRob(i, 0, nums, k)
	}
	res := maxRobList[0]
	for i := 1; i < len(maxRobList); i++ {
		res = min(res, maxRobList[i])
	}
	return res
}

func dfsGetMaxRob(index int, currentMax int, nums []int, k int) {
	if currentMax < nums[index] {
		currentMax = nums[index]
	}
	if k == 1 {
		maxRobList = append(maxRobList, currentMax)
		return
	}
	for i := index + 2; i < len(nums); i++ {
		dfsGetMaxRob(i, currentMax, nums, k-1)
	}
}

func Test_minCapability() {
	// fmt.Println(minCapability([]int{2, 3, 5, 9}, 2))
	fmt.Println(minCapability([]int{25038, 3053, 2825, 3638, 4648, 3259, 4948, 4248, 4940, 2834, 109, 5224, 5097, 4708, 2162, 3438, 4152, 4134, 551, 3961, 2294, 3961, 1327, 2395, 1002, 763, 4296, 3147, 5069, 2156, 572, 1261, 4272, 4158, 5186, 2543, 5055, 4735, 2325, 1206, 1019, 1257, 5048, 1563, 3507, 4269, 5328, 173, 5007, 2392, 967, 2768, 86, 3401, 3667, 4406, 4487, 876, 1530, 819, 1320, 883, 1101, 5317, 2305, 89, 788, 1603, 3456, 5221, 1910, 3343, 4597}, 28))
}
