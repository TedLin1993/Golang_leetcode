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

func TestSequentialDigits() {
	fmt.Println(sequentialDigits(100, 300))
	fmt.Println(sequentialDigits(1000, 13000))
	fmt.Println(sequentialDigits(123, 6789))
	fmt.Println(sequentialDigits(1, 12))
}
