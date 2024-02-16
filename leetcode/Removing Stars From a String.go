package leetcode

import "slices"

func removeStars(s string) string {
	res := make([]byte, 0, len(s))
	starCount := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '*' {
			starCount++
		} else {
			if starCount > 0 {
				starCount--
			} else {
				res = append(res, s[i])
			}
		}
	}
	slices.Reverse(res)
	return string(res)
}

func removeStars_2(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
