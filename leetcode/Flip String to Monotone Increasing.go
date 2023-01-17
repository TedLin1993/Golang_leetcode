package leetcode

import "fmt"

func minFlipsMonoIncr(s string) int {
	dpZero := 0
	dpOne := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			dpOne = min(dpZero, dpOne+1)
		} else {
			dpZero++
		}
	}
	return dpOne
}

func Test_minFlipsMonoIncr() {
	fmt.Println(minFlipsMonoIncr("00110"))      //1
	fmt.Println(minFlipsMonoIncr("010110"))     //2
	fmt.Println(minFlipsMonoIncr("00011000"))   //2
	fmt.Println(minFlipsMonoIncr("0111100000")) //4
}
