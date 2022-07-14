package leetcode

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	f1 := 0
	f2 := 1
	for i := 2; i <= n; i++ {
		temp := f1 + f2
		f1 = f2
		f2 = temp
	}
	return f2
}

func Testfib() {
	fmt.Println(fib(0)) //0
	fmt.Println(fib(1)) //1
	fmt.Println(fib(2)) //1
	fmt.Println(fib(3)) //2
	fmt.Println(fib(4)) //3
}
