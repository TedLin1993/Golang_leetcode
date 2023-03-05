package leetcode

import "fmt"

func passThePillow(n int, time int) int {
	time = time % (2 * (n - 1))
	if time < n {
		return time + 1
	}
	return 2*n - time - 1
}

func Test_passThePillow() {
	fmt.Println(passThePillow(4, 5))  //2
	fmt.Println(passThePillow(3, 2))  //3
	fmt.Println(passThePillow(3, 10)) //3
}
