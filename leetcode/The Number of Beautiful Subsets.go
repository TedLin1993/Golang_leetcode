package leetcode

import "fmt"

func beautifulSubsets(nums []int, k int) int {
	sets := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			diff := nums[j] - nums[i]
			if diff < 0 {
				diff = -diff
			}
			if diff != k {
				sets = append(sets, []int{i, j})
			}
		}
	}

	res := len(nums) + len(sets)
	// for len(sets) > 1 {
	// 	nextSet := make([][]int, 0, len(sets))
	// 	//TODO
	// 	for i := 0; i < len(sets); i++ {
	// 		for j := i + 1; j < len(sets); j++ {
	// 			set1 := sets[i]
	// 			set2 := sets[j]
	// 			count := 0
	// 			for k := 0; k< len(sets[j]); k++ {
	// 				if
	// 			}
	// 		}
	// 	}

	// 	sets = nextSet
	// }

	return res
}

// func canMerge([][]int) (bool, []int) {

// }

func Test_beautifulSubsets() {
	// fmt.Println(beautifulSubsets([]int{2, 4, 6}, 2)) //4
	fmt.Println(beautifulSubsets([]int{2, 4, 6, 7}, 2)) //9
}
