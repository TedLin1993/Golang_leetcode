package leetcode

import "strings"

func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	n := len(s)
	res := s
	left := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			left = i
			break
		}
	}
	cur := 0
	for right := left; right < n; right++ {
		if s[right] == '1' {
			cur++
		}
		if cur > k {
			left++
			cur--
			for s[left] != '1' {
				left++
			}
		}
		if cur == k {
			temp := s[left : right+1]
			if len(temp) < len(res) || (len(temp) == len(res) && temp < res) {
				res = temp
			}
		}
	}
	return res
}
