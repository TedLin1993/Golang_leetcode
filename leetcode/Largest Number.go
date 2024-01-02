package leetcode

import (
	"cmp"
	"fmt"
	"slices"
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

func largestNumber_2(nums []int) string {
	n := len(nums)
	numStrs := make([]string, n)
	for i, num := range nums {
		numStrs[i] = strconv.Itoa(num)
	}
	slices.SortFunc(numStrs, func(a, b string) int {
		s1, s2 := a+b, b+a
		return cmp.Compare(s2, s1)
	})
	if numStrs[0] == "0" {
		return "0"
	}
	return strings.Join(numStrs, "")
}

func Test_largestNumber() {
	fmt.Println(largestNumber([]int{3, 30}))       //330
	fmt.Println(largestNumber([]int{321, 3213}))   //3213321
	fmt.Println(largestNumber([]int{345, 3453}))   //3453453
	fmt.Println(largestNumber([]int{3345, 33453})) //334533453
	fmt.Println(largestNumber([]int{33453, 3345})) //334533453
	fmt.Println(largestNumber([]int{0, 0}))        //0
}
