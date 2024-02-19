package leetcode

import "math"

func longestCommonPrefix_3043(arr1 []int, arr2 []int) int {
	preMap := map[int]bool{}
	for _, v := range arr1 {
		for v > 0 {
			preMap[v] = true
			v /= 10
		}
	}
	res := 0
	for _, v := range arr2 {
		temp := 0
		for v > 0 {
			if preMap[v] {
				temp = int(math.Log10(float64(v))) + 1
				break
			}
			v /= 10
		}
		res = max(res, temp)
	}
	return res
}
