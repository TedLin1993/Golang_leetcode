package leetcode

import (
	"fmt"
	"math/bits"
)

func evenOddBit(n int) []int {
	return []int{bits.OnesCount(uint(n & 0b101010101)), bits.OnesCount(uint(n & 0b10101010))}
}

func Test_evenOddBit() {
	fmt.Println(evenOddBit(17)) //[2, 0]
	fmt.Println(evenOddBit(2))  //[0, 1]
}
