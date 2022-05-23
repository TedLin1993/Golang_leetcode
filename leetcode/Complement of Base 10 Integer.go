package leetcode

import "fmt"

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	res := 0
	power := 1
	for n != 0 {
		if n%2 == 0 {
			res += power
		}
		power = power << 1
		n = n >> 1
	}
	return res
}

func TestComplement_of_Base_10_Integer() {
	fmt.Println(bitwiseComplement(5))
	fmt.Println(bitwiseComplement(7))
	fmt.Println(bitwiseComplement(10))
}
