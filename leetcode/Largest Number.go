package leetcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	n := len(nums)
	numsStr := make([]string, n)
	for i, num := range nums {
		numsStr[i] = strconv.Itoa(num)
	}
	sort.Slice(numsStr, func(a, b int) bool {
		aLen, bLen := len(numsStr[a]), len(numsStr[b])
		for i := 0; i < 10; i++ {
			aIdx, bIdx := i%aLen, i%bLen
			if numsStr[a][aIdx] > numsStr[b][bIdx] {
				return true
			}
			if numsStr[a][aIdx] < numsStr[b][bIdx] {
				return false
			}
		}
		return true
	})
	if numsStr[0] == "0" {
		return "0"
	}
	return strings.Join(numsStr, "")
}

func Test_largestNumber() {
	fmt.Println(largestNumber([]int{3, 30}))       //330
	fmt.Println(largestNumber([]int{321, 3213}))   //3213321
	fmt.Println(largestNumber([]int{345, 3453}))   //3453453
	fmt.Println(largestNumber([]int{3345, 33453})) //334533453
	fmt.Println(largestNumber([]int{33453, 3345})) //334533453
	fmt.Println(largestNumber([]int{0, 0}))        //0
}
