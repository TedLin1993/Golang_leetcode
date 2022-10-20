package leetcode

import (
	"fmt"
	"strings"
)

type RomanInt struct {
	Roman string
	Int   int
}

func intToRoman(num int) string {
	romanInts := romanIntsInitial()
	index := 0
	var res strings.Builder
	for num > 0 {
		if num >= romanInts[index].Int {
			num -= romanInts[index].Int
			res.WriteString(romanInts[index].Roman)
		} else {
			index++
		}
	}
	return res.String()
}

func romanIntsInitial() []RomanInt {
	return []RomanInt{
		{
			Roman: "M",
			Int:   1000,
		},
		{
			Roman: "CM",
			Int:   900,
		},
		{
			Roman: "D",
			Int:   500,
		},
		{
			Roman: "CD",
			Int:   400,
		},
		{
			Roman: "C",
			Int:   100,
		},
		{
			Roman: "XC",
			Int:   90,
		},
		{
			Roman: "L",
			Int:   50,
		},
		{
			Roman: "XL",
			Int:   40,
		},
		{
			Roman: "X",
			Int:   10,
		},
		{
			Roman: "IX",
			Int:   9,
		},
		{
			Roman: "V",
			Int:   5,
		},
		{
			Roman: "IV",
			Int:   4,
		},
		{
			Roman: "I",
			Int:   1,
		},
	}
}

func Test_intToRoman() {
	fmt.Println(intToRoman(3))    //III
	fmt.Println(intToRoman(58))   //LVIII
	fmt.Println(intToRoman(1994)) //MCMXCIV
	fmt.Println(intToRoman(3999)) //MMMCMXCIX
}
