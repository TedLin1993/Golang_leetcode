package leetcode

import "fmt"

func longestValidParentheses(s string) int {
	stack := []int{0}
	leftCount := 0
	rightCount := 0
	for _, char := range s {
		if char == '(' {
			stack = append(stack, 0)
			leftCount++
		} else {
			if leftCount < rightCount+1 {
				//separate valid bracket
				stack = append(stack, 0)
				continue
			}
			rightCount++
			stack[len(stack)-1] += 2
			value := stack[len(stack)-2] + stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], value)
		}
	}

	res := 0
	for _, v := range stack {
		res = max(res, v)
	}
	return res
}

func TestlongestValidParentheses() {
	fmt.Println(longestValidParentheses("(()"))          //2
	fmt.Println(longestValidParentheses("((())"))        //4
	fmt.Println(longestValidParentheses(""))             //0
	fmt.Println(longestValidParentheses(")(())(()()()")) //6
	fmt.Println(longestValidParentheses(")()())"))       //4
	fmt.Println(longestValidParentheses("()(())"))       //6
	fmt.Println(longestValidParentheses(")()())()()("))  //4
}
