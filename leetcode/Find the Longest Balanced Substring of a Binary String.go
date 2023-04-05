package leetcode

import "fmt"

func findTheLongestBalancedSubstring(s string) int {
	oneCount := 0
	res, temp := 0, 0
	preBit := byte('0')
	if s[len(s)-1] == '1' {
		preBit = '1'
		oneCount++
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == preBit {
			if s[i] == '1' {
				oneCount++
			} else {
				if oneCount > 0 {
					oneCount--
					temp += 2
				}
			}
		} else {
			if s[i] == '1' {
				oneCount = 1
				temp = 0
			} else {
				if oneCount > 0 {
					oneCount--
					temp += 2
				}
			}
		}
		preBit = s[i]
		if temp > res {
			res = temp
		}
	}
	return res
}

func Test_findTheLongestBalancedSubstring() {
	fmt.Println(findTheLongestBalancedSubstring("01000111")) //6
	fmt.Println(findTheLongestBalancedSubstring("00111"))    //4
	fmt.Println(findTheLongestBalancedSubstring("111"))      //0
	fmt.Println(findTheLongestBalancedSubstring("0011"))     //4
}
