package leetcode

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	//先區分實數與虛數
	num1f := strings.Split(num1, "+")
	real1, _ := strconv.Atoi(num1f[0])
	img1, _ := strconv.Atoi(strings.Trim(num1f[1], "i"))

	num2f := strings.Split(num2, "+")
	real2, _ := strconv.Atoi(num2f[0])
	img2, _ := strconv.Atoi(strings.Trim(num2f[1], "i"))

	a2 := real1 * real2
	ab := real1 * img2
	ba := img1 * real2
	b2 := img1 * img2

	res := strconv.Itoa(a2-b2) + "+" + strconv.Itoa(ab+ba) + "i"
	return res
}
