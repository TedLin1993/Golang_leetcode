package leetcode

import (
	"fmt"
)

func reorderedPowerOf2(n int) bool {
	digits := countDigits(n)

	// integer limit is 10^9
	// so biggest power of 2 is 2^29
	powerOf2 := 1
	for i := 0; i <= 29; i++ {
		digitsOfPO2 := countDigits(powerOf2)
		if digits == digitsOfPO2 {
			return true
		}
		powerOf2 <<= 1
	}
	return false
}

func countDigits(n int) [10]int {
	digits := [10]int{}
	for n > 0 {
		digits[n%10]++
		n /= 10
	}
	return digits
}

func Test_reorderedPowerOf2() {
	fmt.Println(reorderedPowerOf2(1))   //true
	fmt.Println(reorderedPowerOf2(10))  //false
	fmt.Println(reorderedPowerOf2(61))  //true
	fmt.Println(reorderedPowerOf2(128)) //true
	fmt.Println(reorderedPowerOf2(182)) //true
}
