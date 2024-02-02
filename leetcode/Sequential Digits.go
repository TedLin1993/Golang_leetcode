package leetcode

import "fmt"

func sequentialDigits(low int, high int) []int {
	res := []int{}
	length := countDigitLength(low)
	totalLength := countDigitLength(high)

	for ; length <= totalLength; length++ {
		for firstDigit := 1; firstDigit < 10; firstDigit++ {
			if firstDigit+length > 10 {
				break
			}
			seq := createSequential(firstDigit, length)
			if seq > high {
				break
			}
			if seq >= low {
				res = append(res, seq)
			}
		}
	}
	return res
}

func createSequential(firstDigit int, length int) int {
	digit := firstDigit
	res := 0
	for i := 0; i < length; i++ {
		res = res*10 + digit
		digit++
	}
	return res
}

func countDigitLength(num int) int {
	res := 0
	for num != 0 {
		num = num / 10
		res++
	}
	return res
}

func sequentialDigits_2(low int, high int) []int {
	var res []int
	for len := 2; len <= 9; len++ {
		for i := 1; i <= 9; i++ {
			if i+len-1 > 9 {
				break
			}
			temp := 0
			for j := i; j < i+len; j++ {
				temp = temp*10 + j
			}
			if temp >= low && temp <= high {
				res = append(res, temp)
			}
			if temp > high {
				return res
			}
		}
	}
	return res
}

func TestSequentialDigits() {
	fmt.Println(sequentialDigits(100, 300))
	fmt.Println(sequentialDigits(1000, 13000))
	fmt.Println(sequentialDigits(123, 6789))
	fmt.Println(sequentialDigits(1, 12))
}
