package leetcode

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	isPositive := false
	if (dividend > 0) == (divisor > 0) {
		isPositive = true
	}
	dividend = abs(dividend)
	divisor = abs(divisor)

	res := 0
	for dividend >= divisor {
		temp, i := divisor, 1
		for dividend >= temp {
			dividend -= temp
			res += i

			temp <<= 1
			i <<= 1
		}

	}

	if !isPositive {
		return -res
	}
	return res
}

func Testdivide() {
	fmt.Println(divide(10, 3))           //3
	fmt.Println(divide(7, -3))           //-2
	fmt.Println(divide(-2147483648, -1)) //2147483647
}
