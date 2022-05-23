package leetcode

import (
	"fmt"
)

func scoreOfParentheses(s string) int {
	stack := []int{0}
	for _, bucket := range s {
		if bucket == '(' {
			stack = append(stack, 0)
		} else {
			valueInsideBucket := stack[len(stack)-1]
			value := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			value = value + max((valueInsideBucket<<1), 1)
			stack = append(stack, value)
		}
	}

	return stack[0]
}

func TestscoreOfParentheses() {
	fmt.Println(scoreOfParentheses("()"))       //1
	fmt.Println(scoreOfParentheses("(())"))     //2
	fmt.Println(scoreOfParentheses("()()"))     //2
	fmt.Println(scoreOfParentheses("(()())"))   //4
	fmt.Println(scoreOfParentheses("(()())()")) //5
}
