package leetcode

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		//負數 % 10 會是負數
		lastElement := x % 10
		//result > 2^31-1
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && lastElement > 7) {
			return 0
		}
		//result < -2^31
		if res < math.MinInt32/10 || (res == math.MaxInt32/10 && lastElement < -8) {
			return 0
		}
		res = 10*res + lastElement
		x = x / 10
	}
	return res
}

func reverse_2(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return res
}

func Test_reverse() {
	var x int
	x = 123
	fmt.Println(reverse(x)) // 321
	x = -123
	fmt.Println(reverse(x)) //-321
	x = 120
	fmt.Println(reverse(x)) //21
	x = 0
	fmt.Println(reverse(x)) //0
	x = 1534236469
	fmt.Println(reverse(x)) //0
}
