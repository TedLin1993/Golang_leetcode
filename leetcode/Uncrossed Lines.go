package leetcode

import "fmt"

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	nums2Map := make([][]int, 2001)
	for idx, num := range nums2 {
		nums2Map[num] = append(nums2Map[num], idx)
	}
	memory := make([][]int, len(nums1))
	for i := 0; i < len(memory); i++ {
		memory[i] = make([]int, len(nums2))
		for j := 0; j < len(memory[i]); j++ {
			memory[i][j] = -1
		}
	}
	var dfs func(nums1Idx, nums2Idx int) int
	dfs = func(nums1Idx, nums2Idx int) int {
		if nums1Idx >= len(nums1) || nums2Idx >= len(nums2) {
			return 0
		}
		if memory[nums1Idx][nums2Idx] != -1 {
			return memory[nums1Idx][nums2Idx]
		}
		num := nums1[nums1Idx]
		nextIdx := -1
		for i := 0; i < len(nums2Map[num]); i++ {
			if nums2Map[num][i] >= nums2Idx {
				nextIdx = nums2Map[num][i] + 1
				break
			}
		}
		res := 0
		if nextIdx > -1 {
			res = 1 + dfs(nums1Idx+1, nextIdx)
		}
		res = max(res, dfs(nums1Idx+1, nums2Idx))
		memory[nums1Idx][nums2Idx] = res
		return res
	}
	return dfs(0, 0)
}

func Test_maxUncrossedLines() {
	fmt.Println(maxUncrossedLines([]int{15, 14, 1, 7, 15, 1, 12, 18, 9, 15, 1, 20, 18, 15, 16, 18, 11, 8, 11, 18, 11, 11, 17, 20, 16, 20, 15, 15, 9, 18, 16, 4, 16, 1, 13, 10, 10, 20, 4, 18, 17, 3, 8, 1, 8, 19, 14, 10, 10, 12}, []int{12, 8, 17, 4, 2, 18, 16, 10, 11, 12, 7, 1, 8, 16, 4, 14, 12, 18, 18, 19, 19, 1, 11, 18, 1, 6, 12, 17, 6, 19, 10, 5, 11, 16, 6, 17, 12, 1, 9, 3, 19, 2, 18, 18, 2, 4, 11, 11, 14, 9, 20, 19, 2, 20, 9, 15, 8, 7, 8, 6, 19, 12, 4, 11, 18, 18, 1, 6, 9, 17, 13, 19, 5, 4, 14, 9, 11, 15, 2, 5, 4, 1, 10, 11, 6, 4, 9, 7, 11, 7, 3, 8, 11, 12, 4, 19, 12, 17, 14, 18}))
}
