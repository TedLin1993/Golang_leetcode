package leetcode

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		//get bits need to carry
		carry := a & b
		//get bits don't need to carry
		a = a ^ b
		//carry bits left shift 1 bit
		b = carry << 1
	}
	return a
}

func TestgetSum() {
	// fmt.Println(getSum(2, 1)) //3
	// fmt.Println(getSum(2, 3)) //5
	fmt.Println(getSum(-1, 1)) //0
}
