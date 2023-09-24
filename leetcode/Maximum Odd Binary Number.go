package leetcode

func maximumOddBinaryNumber(s string) string {
	n := len(s)
	oneCount := 0
	for _, b := range s {
		if b == '1' {
			oneCount++
		}
	}
	res := ""
	for i := 0; i < oneCount-1; i++ {
		res += "1"
	}
	for i := 0; i < n-oneCount; i++ {
		res += "0"
	}
	res += "1"
	return res
}
