package leetcode

import (
	"fmt"
	"strconv"
)

func maximum69Number(num int) int {
	strNum := strconv.Itoa(num)
	for i := 0; i < len(strNum); i++ {
		if strNum[i] == '6' {
			strNum = strNum[:i] + "9" + strNum[i+1:]
			break
		}
	}
	v, _ := strconv.Atoi(strNum)
	return v
}

func Test_maximum69Number() {
	fmt.Println(maximum69Number(9669)) //9969
	fmt.Println(maximum69Number(9996)) //9999
	fmt.Println(maximum69Number(9999)) //9999
}
