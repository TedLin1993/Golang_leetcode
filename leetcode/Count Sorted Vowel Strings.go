package leetcode

import "fmt"

func countVowelStrings(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	for n > 1 {
		a = a + e + i + o + u
		e = e + i + o + u
		i = i + o + u
		o = o + u
		u = u
		n--
	}
	return a + e + i + o + u
}

func TestcountVowelStrings() {
	fmt.Println(countVowelStrings(1))  // 5
	fmt.Println(countVowelStrings(2))  // 15
	fmt.Println(countVowelStrings(33)) // 66045
	fmt.Println(countVowelStrings(50)) // 316251
}
