package leetcode

import (
	"fmt"
)

func minimumString(a string, b string, c string) string {
	slices := [][]string{{a, b, c}, {a, c, b}, {b, a, c}, {b, c, a}, {c, a, b}, {c, b, a}}
	res := a + b + c
	for _, slice := range slices {
		temp := slice[0]
		for _, s := range slice[1:] {
			temp = mergeString(temp, s)
		}
		if len(temp) < len(res) || (len(temp) == len(res) && temp < res) {
			res = temp
		}
	}

	return res
}

func mergeString(a, b string) string {
	n, m := len(a), len(b)
	isContain := false
	for i := 0; i < n; i++ {
		if a[i] == b[0] {
			if n-i >= m && a[i:i+m] == b {
				isContain = true
				break
			} else if n-i < m && a[i:] == b[:n-i] {
				isContain = true
				a = a + b[n-i:]
				break
			}
		}
	}
	if !isContain {
		a += b
	}
	return a
}

func Test_minimumString() {
	fmt.Println(minimumString("abc", "bca", "aaa")) //aaabca
	fmt.Println(minimumString("ab", "ba", "aba"))   //aba
	fmt.Println(minimumString("a", "a", "ca"))      //ca
	fmt.Println(minimumString("b", "a", "ba"))      //ba
	fmt.Println(minimumString("ab", "bcd", "cd"))   //abcd
	fmt.Println(minimumString("a", "cc", "a"))      //acc
	fmt.Println(minimumString("d", "j", "oa"))      //djoa
	fmt.Println(minimumString("b", "c", "aa"))      //aabc
	fmt.Println(minimumString("ac", "c", "b"))      //acb
}
