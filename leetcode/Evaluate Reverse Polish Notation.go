package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, t := range tokens {
		if t == "+" || t == "-" || t == "*" || t == "/" {
			temp := 0
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			switch t {
			case "+":
				temp = a + b
			case "-":
				temp = a - b
			case "*":
				temp = a * b
			case "/":
				temp = a / b
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, temp)
		} else {
			v, _ := strconv.Atoi(t)
			stack = append(stack, v)
		}
	}
	return stack[0]
}
