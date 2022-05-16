package leetcode

import (
	"fmt"
)

type decodeInfo struct {
	times int
	pos   int
}

func decodeString(s string) string {
	stack := []decodeInfo{}
	times := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			times = times*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			stack = append(stack, decodeInfo{times: times, pos: i})
			times = 0
		} else if s[i] == ']' {
			df := pop(stack)
			stack = stack[:len(stack)-1]
			decodeStr := reapeat(s[df.pos+1:i], df.times)
			digitCount := recursionCountDigits(df.times)
			s = s[:df.pos-digitCount] + decodeStr + s[i+1:]
			i = df.pos - digitCount + len(decodeStr) - 1
		}
	}
	return s
}

func pop(stack []decodeInfo) decodeInfo {
	res := stack[len(stack)-1]
	return res
}

func reapeat(s string, times int) string {
	res := ""
	for i := 0; i < times; i++ {
		res = res + s
	}
	return res
}

func recursionCountDigits(number int) int {
	if number < 10 {
		return 1
	} else {
		return 1 + recursionCountDigits(number/10)
	}
}

func TestDecodeString() {
	fmt.Println(decodeString("3[a]2[bc]"))     //aaabcbc
	fmt.Println(decodeString("3[a2[c]]"))      //accaccacc
	fmt.Println(decodeString("2[abc]3[cd]ef")) //abcabccdcdcdef
	fmt.Println(decodeString("abc3[cd]xyz"))   //abccdcdcdxyz
	fmt.Println(decodeString("10[a]2[bc]"))    //aaaaaaaaaabcbc

}
