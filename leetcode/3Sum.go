package leetcode

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)

	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			}
		}

	}
	return res
}

func Test_threeSum() {
	var nums []int
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	nums = []int{}
	fmt.Println(threeSum(nums))
	nums = []int{0}
	fmt.Println(threeSum(nums))
	nums = []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}
