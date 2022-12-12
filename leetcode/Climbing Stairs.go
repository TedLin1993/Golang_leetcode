package leetcode

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		temp := a + b
		a, b = b, temp
	}
	return b
}

func Test_climbStairs() {
	fmt.Println(climbStairs(1))  //1
	fmt.Println(climbStairs(2))  //2
	fmt.Println(climbStairs(3))  //3
	fmt.Println(climbStairs(10)) //89
	fmt.Println(climbStairs(45)) //1836311903
}
