package leetcode

import "fmt"

func subarraysDivByK(nums []int, k int) int {
	res := 0
	sum := 0
	modMap := make([]int, k)
	modMap[0] = 1
	for _, num := range nums {
		sum += num
		mod := sum % k
		if mod < 0 {
			mod += k
		}
		res += modMap[mod]
		modMap[mod]++
	}
	return res
}

func Test_subarraysDivByK() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))                                         //7
	fmt.Println(subarraysDivByK([]int{5}, 9))                                                          //0
	fmt.Println(subarraysDivByK([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 9)) //210
}
