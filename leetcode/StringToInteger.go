package leetcode

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}

	isNegtive := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		isNegtive = true
		s = s[1:]
	}

	res := 0
	for i := 0; i < len(s); i++ {
		if !isDigit(s[i]) {
			break
		}
		if isOverLimit(res, s[i], isNegtive) {
			if isNegtive {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		res = res*10 + int(s[i]-'0')
	}

	if isNegtive {
		res = -res
	}

	return res
}

func isDigit(char byte) bool {
	if char < '0' || char > '9' {
		return false
	}
	return true
}

func isOverLimit(num int, digit byte, isNegtive bool) bool {
	if num > 214748364 {
		return true
	}
	if num == 214748364 {
		if !isNegtive && digit-'0' > 7 {
			return true
		}
		if isNegtive && digit-'0' > 8 {
			return true
		}
	}
	return false
}

func TestStringToInteger() {
	fmt.Println(myAtoi("42"))              //42
	fmt.Println(myAtoi("   -42"))          //-42
	fmt.Println(myAtoi("4193 with words")) //4193
	fmt.Println(myAtoi("words and 987"))   //0
	fmt.Println(myAtoi("+-12"))            //0
}
