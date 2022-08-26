package leetcode

import (
	"fmt"
)

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	//upper bound is 3^19 = 1162261467
	return 1162261467%n == 0
}

func Test_isPowerOfThree() {
	fmt.Println(isPowerOfThree(1))        //true
	fmt.Println(isPowerOfThree(3))        //true
	fmt.Println(isPowerOfThree(0))        //false
	fmt.Println(isPowerOfThree(27))       //true
	fmt.Println(isPowerOfThree(43046720)) //false
	fmt.Println(isPowerOfThree(-3))       //false
}
