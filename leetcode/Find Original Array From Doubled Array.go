package leetcode

import (
	"fmt"
	"sort"
)

func findOriginalArray(changed []int) []int {
	arrMap := map[int]int{}
	for _, v := range changed {
		arrMap[v]++
	}
	sort.Ints(changed)
	res := []int{}
	for _, v := range changed {
		if arrMap[v] == 0 {
			continue
		}
		arrMap[v]--
		if arrMap[v*2] > 0 {
			res = append(res, v)
			arrMap[v*2]--
		} else {
			return nil
		}
	}
	return res
}

func Test_findOriginalArray() {
	fmt.Println(findOriginalArray([]int{1, 3, 4, 2, 6, 8}))          //[1 3 4]
	fmt.Println(findOriginalArray([]int{6, 3, 0, 1}))                //[]
	fmt.Println(findOriginalArray([]int{0}))                         //[]
	fmt.Println(findOriginalArray([]int{0, 0, 0, 0}))                //[0 0]
	fmt.Println(findOriginalArray([]int{4, 4, 16, 20, 8, 8, 2, 10})) //[2 4 8 10]

}
