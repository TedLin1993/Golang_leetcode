package leetcode

import (
	"fmt"
)

func countAndSay(n int) string {
	//When n=1, result is 1. This is base case.
	res := []byte{'1'}

	for i := 1; i < n; i++ {
		next := []byte{}
		digit := res[0]
		count := byte('1')
		index := 1
		for index < len(res) {
			if res[index] == digit {
				count++
			} else {
				next = append(next, byte(count), digit)
				digit = res[index]
				count = byte('1')
			}
			index++
		}
		next = append(next, byte(count), digit)
		res = next
	}
	return string(res)
}

func Test_countAndSay() {
	fmt.Println(countAndSay(1))  //1
	fmt.Println(countAndSay(4))  //1211
	fmt.Println(countAndSay(10)) //13211311123113112211
}
