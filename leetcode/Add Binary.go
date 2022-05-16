package leetcode

import "fmt"

func addBinary(a string, b string) string {
	res := ""
	aIndex := len(a) - 1
	bIndex := len(b) - 1
	var carry byte = 0

	for aIndex >= 0 && bIndex >= 0 {
		sum := (a[aIndex] - '0') + (b[bIndex] - '0') + carry
		carry, res = getCarry(sum, res)
		aIndex--
		bIndex--
	}
	for aIndex >= 0 {
		sum := (a[aIndex] - '0') + carry
		carry, res = getCarry(sum, res)
		aIndex--
	}
	for bIndex >= 0 {
		sum := (b[bIndex] - '0') + carry
		carry, res = getCarry(sum, res)
		bIndex--
	}
	if carry > 0 {
		res = "1" + res
	}

	return res
}

func getCarry(sum byte, res string) (byte, string) {
	var carry byte
	switch sum {
	case 3:
		carry = 1
		res = "1" + res
	case 2:
		carry = 1
		res = "0" + res
	case 1:
		carry = 0
		res = "1" + res
	case 0:
		carry = 0
		res = "0" + res
	}
	return carry, res
}

func TestAdd_Binary() {
	// fmt.Println(addBinary("11", "1"))
	// fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("1", "111"))
}
