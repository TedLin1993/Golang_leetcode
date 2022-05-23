package leetcode

import "fmt"

func maxPower(s string) int {
	if len(s) < 2 {
		return 1
	}
	index := 0
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			res = max(res, i-index)
			index = i
		}
	}
	res = max(res, len(s)-index)
	return res
}

func TestConsecutiveCharacters() {
	fmt.Println(maxPower("leetcode"))            //2
	fmt.Println(maxPower("abbcccddddeeeeedcba")) //5
	fmt.Println(maxPower("triplepillooooow"))    //5
	fmt.Println(maxPower("hooraaaaaaaaaaay"))    //11
	fmt.Println(maxPower("tourist"))             //1
	fmt.Println(maxPower("cc"))                  //2
}
