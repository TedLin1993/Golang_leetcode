package leetcode

import "fmt"

func romanToInt(s string) int {
	romanIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := romanIntMap[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if romanIntMap[s[i]] >= romanIntMap[s[i+1]] {
			res += romanIntMap[s[i]]
		} else {
			res -= romanIntMap[s[i]]
		}
	}
	return res
}

func Test_romanToInt() {
	var s string
	s = "III"
	fmt.Println(romanToInt(s)) //3
	s = "IV"
	fmt.Println(romanToInt(s)) //4
	s = "IX"
	fmt.Println(romanToInt(s)) //9
	s = "LVIII"
	fmt.Println(romanToInt(s)) //58
	s = "MCMXCIV"
	fmt.Println(romanToInt(s)) //1994
}
