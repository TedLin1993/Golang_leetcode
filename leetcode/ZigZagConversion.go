package leetcode

import "fmt"

//若 numRows = 3
//第一排是 0,4,8,12,...4n
//第二排是 1,3,5,7,...4n+1和4n-1交錯
//第三排是 2,6,10,...4n+2

//若 numRows = 4
// 0,6,12,... 6n
// 1,5,7,11 ... 6n+1 and 6n+5
// 2,4,8,10 ... 6n+2 and 6n+4
// 3,9,15,... 6n+3

// n = round index
// row0:(2numRows-2)n
// row1:(2numRows-2)n+1 and (2numRows-2)n+(2numRows-3)
// rowk:(2numRows-2)n+k and (2numRows-2)n+(2numRows-2-k)
// last row: (2numRows-2)n+(2numRows-1)
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var res string
	//zigzag count
	roundCount := len(s) / (2*numRows - 2)
	//last items after zigzag
	remainingCount := len(s) % (2*numRows - 2)
	for k := 0; k < numRows; k++ {
		for i := 0; i < roundCount; i++ {
			res += string(s[(2*numRows-2)*i+k])
			if k != 0 && k != numRows-1 {
				res += string(s[(2*numRows-2)*i+(2*numRows-2-k)])
			}
		}
		if remainingCount > k {
			res += string(s[(2*numRows-2)*roundCount+k])
		}
		if k != 0 && k != numRows-1 && remainingCount > 2*numRows-2-k {
			res += string(s[(2*numRows-2)*roundCount+(2*numRows-2-k)])
		}
	}

	return res
}

func convert_2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	charLists := make([][]byte, numRows)
	idx := 0
	isIncrease := true
	for i := 0; i < len(s); i++ {
		charLists[idx] = append(charLists[idx], s[i])
		updateIdx(&idx, &isIncrease, numRows)
	}
	res := make([]byte, 0, len(s))
	for _, charList := range charLists {
		res = append(res, charList...)
	}
	return string(res)
}

func updateIdx(idx *int, isIncrease *bool, numRows int) {
	if *idx == 0 {
		*isIncrease = true
	} else if *idx == numRows-1 {
		*isIncrease = false
	}
	if *isIncrease {
		*idx++
	} else {
		*idx--
	}
}

func Test_convert() {
	var s string
	var numRows int

	s = "PAYPALISHIRING"
	numRows = 3
	fmt.Println(convert(s, numRows) == "PAHNAPLSIIGYIR")

	s = "PAYPALISHIRING"
	numRows = 4
	fmt.Println(convert(s, numRows) == "PINALSIGYAHRPI")

	s = "PAYPALISHIRI"
	numRows = 3
	fmt.Println(convert(s, numRows) == "PAHAPLSIIYIR")

	s = "A"
	numRows = 1
	fmt.Println(convert(s, numRows) == "A")

	s = "AB"
	numRows = 3
	fmt.Println(convert(s, numRows) == "AB")
}
