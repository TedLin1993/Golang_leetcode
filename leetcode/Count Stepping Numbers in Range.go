package leetcode

import "fmt"

var mod = int(1e9) + 7

func countSteppingNumbers(low string, high string) int {
	res := (calc(high) - calc(low) + mod) % mod
	isValid := func(s string) bool {
		for i := 1; i < len(s); i++ {
			if abs(int(s[i])-int(s[i-1])) != 1 {
				return false
			}
		}
		return true
	}
	if isValid(low) {
		res = (res + 1) % mod
	}
	return res
}

func calc(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, 10)
	}
	var count func(pos int, preDigit int, isLimit bool, isNum bool) int
	count = func(pos int, preDigit int, isLimit bool, isNum bool) int {
		if pos == len(s) {
			if isNum {
				return 1
			}
			return 0
		}
		if !isLimit && isNum && dp[pos][preDigit] != 0 {
			return dp[pos][preDigit]
		}

		res := 0
		if !isNum {
			res = (res + count(pos+1, 0, false, false)) % mod
		}
		left := 0
		if !isNum {
			left = 1
		}
		right := 9
		if isLimit {
			right = int(s[pos] - '0')
		}
		for digit := left; digit <= right; digit++ {
			if !isNum || abs(digit-preDigit) == 1 {
				res = (res + count(pos+1, digit, isLimit && digit == right, true)) % mod
			}
		}
		if !isLimit && isNum {
			dp[pos][preDigit] = res
		}
		return res
	}
	res := count(0, 0, true, false)
	return res
}

func Test_countSteppingNumbers() {
	// fmt.Println(countSteppingNumbers("1", "11"))   //10
	// fmt.Println(countSteppingNumbers("90", "101")) //2
	fmt.Println(countSteppingNumbers("1", "100000000000000")) //84408
}
