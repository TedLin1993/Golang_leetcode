package leetcode

import "fmt"

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}
	return res
}

func TesthammingWeight() {
	fmt.Println(hammingWeight(00000000000000000000000000001011)) //3
	fmt.Println(hammingWeight(00000000000000000000000010000000)) //1
	// fmt.Println(hammingWeight(11111111111111111111111111111101)) //31
}
