package leetcode

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right := 1, x/2
	for left < right {
		mid := left + (right-left+1)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}

func Test_mySqrt() {
	fmt.Println(mySqrt(0)) //0
	fmt.Println(mySqrt(2)) //1
	fmt.Println(mySqrt(4)) //2
	fmt.Println(mySqrt(8)) //2
}
