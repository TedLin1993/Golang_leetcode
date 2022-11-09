package leetcode

import "fmt"

var charDist byte = 'a' - 'A'

func makeGood(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || !isBadAdjacent(stack[len(stack)-1], s[i]) {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

func isBadAdjacent(a, b byte) bool {
	if a-charDist == b || a+charDist == b {
		return true
	}
	return false
}

func Test_makeGood() {
	fmt.Println(makeGood("leEeetcode")) //leetcode
	fmt.Println(makeGood("abBAcC"))     //
	fmt.Println(makeGood("s"))          //s
}
