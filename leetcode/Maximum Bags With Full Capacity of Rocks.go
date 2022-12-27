package leetcode

import (
	"fmt"
	"sort"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	fullBagCount := 0
	remainCapacities := make([]int, 0, len(capacity))
	for i := 0; i < len(capacity); i++ {
		remainCapacity := capacity[i] - rocks[i]
		if remainCapacity == 0 {
			fullBagCount++
		} else {
			remainCapacities = append(remainCapacities, remainCapacity)
		}
	}
	sort.Ints(remainCapacities)
	for _, remain := range remainCapacities {
		if remain <= additionalRocks {
			additionalRocks -= remain
			fullBagCount++
		} else {
			break
		}
	}
	return fullBagCount
}

func Test_maximumBags() {
	fmt.Println(maximumBags([]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2)) //3
	fmt.Println(maximumBags([]int{10, 2, 2}, []int{2, 2, 0}, 100))    //3
}
