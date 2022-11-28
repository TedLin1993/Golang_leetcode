package leetcode

import "fmt"

var modulo int = 1000000007

func sumSubarrayMins(arr []int) int {
	res := 0
	stack := []int{-1}
	for i := 0; i < len(arr); i++ {
		last := stack[len(stack)-1]
		for last != -1 && arr[last] >= arr[i] {
			add := getSubarraySum(&stack, arr, i)
			res = (res + add) % modulo
			last = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 1 {
		add := getSubarraySum(&stack, arr, len(arr))
		res = (res + add) % modulo
	}
	return res
}

func getSubarraySum(stack *[]int, arr []int, index int) int {
	mid := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	rightCount := index - mid
	leftCount := mid - (*stack)[len(*stack)-1]
	return leftCount * rightCount * arr[mid] % modulo

}

func Test_sumSubarrayMins() {
	fmt.Println(sumSubarrayMins([]int{3}))                 //3
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))        //17
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 2, 4}))     //25
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3})) //444
}
