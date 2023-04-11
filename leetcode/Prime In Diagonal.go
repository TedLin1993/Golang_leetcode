package leetcode

import (
	"fmt"
)

func diagonalPrime(nums [][]int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i][i] > res && isPrime(nums[i][i]) {
			res = nums[i][i]
		}
		if nums[i][n-i-1] > res && isPrime(nums[i][n-i-1]) {
			res = nums[i][n-i-1]
		}
	}
	return res
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Test_diagonalPrime() {
	fmt.Println(diagonalPrime([][]int{{1, 2, 3}, {5, 17, 7}, {9, 11, 20}})) //17
}
