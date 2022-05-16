package leetcode

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func TestisPowerOfTwo() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
}
