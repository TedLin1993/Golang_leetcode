package leetcode

import (
	"fmt"
)

func calculate(s string) int {
	stack := []int{}
	value := 0
	operator := byte('+')
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			value = value*10 + int(s[i]-'0')
		} else {
			pushStack(&stack, operator, &value)
			operator = s[i]
		}
	}
	//push last value to stack
	pushStack(&stack, operator, &value)

	res := 0
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}

func pushStack(stack *[]int, operator byte, value *int) {
	switch operator {
	case '*':
		(*stack)[len((*stack))-1] = (*stack)[len((*stack))-1] * *value
	case '/':
		(*stack)[len((*stack))-1] = (*stack)[len((*stack))-1] / *value
	case '+':
		(*stack) = append((*stack), *value)
	case '-':
		(*stack) = append((*stack), -*value)
	}
	*value = 0
}

func TestBasic_Calculator_II() {
	fmt.Println(calculate("1"))           //1
	fmt.Println(calculate("1 + 1"))       //2
	fmt.Println(calculate("3+2*2"))       //7
	fmt.Println(calculate(" 3/2 "))       //1
	fmt.Println(calculate(" 3+5 / 2 "))   //5
	fmt.Println(calculate(" 13+25 / 5 ")) //18
	fmt.Println(calculate("1 - 1"))       //0
	fmt.Println(calculate("2*3+4"))       //10

}
