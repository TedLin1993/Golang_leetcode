package leetcode

import "fmt"

func findValidSplit(nums []int) int {
	suffixProduct := 1
	for _, num := range nums {
		suffixProduct *= num
	}
	prefixProduct := 1
	for i := 0; i < len(nums)-1; i++ {
		suffixProduct /= nums[i]
		prefixProduct *= nums[i]
		gcd := gcd(suffixProduct, prefixProduct)
		if gcd == 1 {
			return i
		}
	}
	return -1
}

func gcd(a, b int) int {
	if a < b {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func Test_findValidSplit() {
	fmt.Println(findValidSplit([]int{4, 7, 8, 15, 3, 5})) //2
	fmt.Println(findValidSplit([]int{4, 7, 15, 8, 3, 5})) //-1
}
