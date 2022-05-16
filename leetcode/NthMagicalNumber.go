package leetcode

import "fmt"

func nthMagicalNumber(n int, a int, b int) int {
	mod := 1000000007
	lcd := getLCD(a, b)
	l := min(a, b)
	r := n * min(a, b)
	getMagicalNumber := func(x int) int { return x/a + x/b - x/lcd }
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if getMagicalNumber(mid) < n {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	//if multiple value has the same magical number, then the smallest one(aka l) is the answer
	return l % mod
}

//最大公因數
func getGCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

//最小公倍數
func getLCD(a, b int) int {
	gcd := getGCD(a, b)
	return a * b / gcd
}

func TestNthMagicalNumber() {
	fmt.Println(nthMagicalNumber(1, 2, 3))                 //2
	fmt.Println(nthMagicalNumber(4, 2, 3))                 //6
	fmt.Println(nthMagicalNumber(5, 2, 4))                 //10
	fmt.Println(nthMagicalNumber(3, 6, 4))                 //8
	fmt.Println(nthMagicalNumber(788907919, 31116, 6062))  //794991020
	fmt.Println(nthMagicalNumber(887859796, 29911, 37516)) //257511204
}
