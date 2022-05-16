package leetcode

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	dictDigitLetter := make(map[byte][]string)
	dictDigitLetter['2'] = []string{"a", "b", "c"}
	dictDigitLetter['3'] = []string{"d", "e", "f"}
	dictDigitLetter['4'] = []string{"g", "h", "i"}
	dictDigitLetter['5'] = []string{"j", "k", "l"}
	dictDigitLetter['6'] = []string{"m", "n", "o"}
	dictDigitLetter['7'] = []string{"p", "q", "r", "s"}
	dictDigitLetter['8'] = []string{"t", "u", "v"}
	dictDigitLetter['9'] = []string{"w", "x", "y", "z"}

	res := []string{}
	for i := 0; i < len(digits); i++ {
		temp := []string{}
		for j := 0; j < len(dictDigitLetter[digits[i]]); j++ {
			digit := dictDigitLetter[digits[i]][j]
			if len(res) == 0 {
				temp = append(temp, digit)
			} else {
				for k := 0; k < len(res); k++ {
					temp = append(temp, res[k]+digit)
				}
			}
		}
		res = temp
	}
	return res
}
