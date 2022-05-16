package leetcode

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(a, b int) bool {
		return people[a] > people[b]
	})
	left := 0
	right := len(people) - 1
	boatCount := 0
	for left <= right {
		remain := limit - people[left]
		left++
		if people[right] <= remain {
			right--
		}
		boatCount++
	}
	return boatCount
}

func TestnumRescueBoats() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))          //1
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))    //3
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))    //4
	fmt.Println(numRescueBoats([]int{5, 1, 4, 2}, 6))    //2
	fmt.Println(numRescueBoats([]int{3, 2, 3, 2, 2}, 6)) //3
}
