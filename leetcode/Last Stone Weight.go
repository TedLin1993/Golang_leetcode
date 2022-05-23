package leetcode

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	sort.Slice(stones, func(i, j int) bool {
		return stones[i] > stones[j]
	})
	for len(stones) > 1 {
		if stones[0] == stones[1] {
			stones = stones[2:]
		} else {
			stones[1] = stones[0] - stones[1]
			stones = stones[1:]
			for i := 0; i < len(stones)-1; i++ {
				if stones[i] < stones[i+1] {
					stones[i], stones[i+1] = stones[i+1], stones[i]
				} else {
					break
				}
			}
		}
	}

	if len(stones) == 1 {
		return stones[0]
	} else {
		return 0
	}
}

func TestlastStoneWeight() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1})) //1
	fmt.Println(lastStoneWeight([]int{1}))                //1
	fmt.Println(lastStoneWeight([]int{2, 2}))             //0
}
