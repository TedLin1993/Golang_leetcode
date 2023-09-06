package leetcode

func countSymmetricIntegers(low int, high int) int {
	res := 0
	for num := low; num <= high; num++ {
		digitCount := getDigitCount(num)
		if digitCount%2 != 0 {
			continue
		}
		if isSymmetric(num, digitCount) {
			res++
		}
	}
	return res
}

func getDigitCount(num int) int {
	res := 0
	for num > 0 {
		num /= 10
		res++
	}
	return res
}

func isSymmetric(num int, digitCount int) bool {
	left := 0
	for i := 0; i < digitCount/2; i++ {
		left += num % 10
		num /= 10
	}
	right := 0
	for num > 0 {
		right += num % 10
		num /= 10
	}
	return left == right
}
