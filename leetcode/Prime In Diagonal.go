package leetcode

import (
	"fmt"
)

func diagonalPrime(nums [][]int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || j == n-i-1 {
				if isPrime(nums[i][j]) && nums[i][j] > res {
					res = nums[i][j]
				}
			}
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
