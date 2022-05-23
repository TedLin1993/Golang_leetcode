package leetcode

import "fmt"

func findComplement(num int) int {
	result := 0
	layer := 1
	for num != 0 {
		if num%2 == 0 {
			result += layer
		}
		layer = layer << 1
		num = num >> 1
	}
	return result
}

func TestNumber_Complement() {
	fmt.Println(findComplement(5)) //2
	fmt.Println(findComplement(1)) //0
	fmt.Println(findComplement(6)) //1
}
