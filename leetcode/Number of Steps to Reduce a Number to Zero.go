package leetcode

import "fmt"

func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			num -= 1
		} else {
			num = num >> 1
		}
		count++
	}
	return count
}

func TestnumberOfSteps() {
	fmt.Println(numberOfSteps(14))  //6
	fmt.Println(numberOfSteps(8))   //4
	fmt.Println(numberOfSteps(123)) //12
}
