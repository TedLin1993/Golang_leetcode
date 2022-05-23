package leetcode

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	popIndex := 0
	for pushIndex := range pushed {
		stack = append(stack, pushed[pushIndex])
		for len(stack) > 0 && stack[len(stack)-1] == popped[popIndex] {
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}
	return len(stack) == 0
}

func TestvalidateStackSequences() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})) //true
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2})) //false
}
