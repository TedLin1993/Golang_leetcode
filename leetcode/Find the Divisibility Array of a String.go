package leetcode

import (
	"fmt"
)

func divisibilityArray(word string, m int) []int {
	res := make([]int, len(word))
	value := 0
	for i := 0; i < len(word); i++ {
		value = (value*10 + int(word[i]-'0')) % m
		if value == 0 {
			res[i] = 1
		}
	}
	return res
}

func Test_divisibilityArray() {
	fmt.Println(divisibilityArray("998244353", 3))
}
