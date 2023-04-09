package leetcode

import (
	"fmt"
	"math"
)

func diagonalPrime(nums [][]int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if (i == j || j == len(nums)-i-1) && isPrime(nums[i][j]) {
				res = max(res, nums[i][j])
			}
		}
	}
	return res
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Test_diagonalPrime() {
	fmt.Println(diagonalPrime([][]int{{1, 2, 3}, {5, 17, 7}, {9, 11, 20}})) //17
}
