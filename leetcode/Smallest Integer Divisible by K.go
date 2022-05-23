package leetcode

import "fmt"

func smallestRepunitDivByK(k int) int {
	reminder := 1
	length := 1
	for i := 0; i < k; i++ {
		if reminder%k == 0 {
			return length
		}
		reminder = 10*(reminder%k) + 1
		length++
	}
	return -1
}

func TestSmallest_Integer_Divisible_by_K() {
	fmt.Println(smallestRepunitDivByK(1)) //1
	fmt.Println(smallestRepunitDivByK(2)) //-1
	fmt.Println(smallestRepunitDivByK(3)) //3
}
