package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func substringXorQueries(s string, queries [][]int) [][]int {
	res := make([][]int, 0, len(queries))
	for _, query := range queries {
		value := query[0] ^ query[1]
		valueString := strconv.FormatInt(int64(value), 2)
		idx := strings.Index(s, valueString)
		if idx == -1 {
			res = append(res, []int{-1, -1})
		} else {
			res = append(res, []int{idx, idx + len(valueString) - 1})
		}
	}
	return res
}

func Test_substringXorQueries() {
	fmt.Println(substringXorQueries("101101", [][]int{{0, 5}, {1, 2}})) //[[0 2] [2 3]]
	fmt.Println(substringXorQueries("0101", [][]int{{12, 8}}))          //[[-1 -1]
	fmt.Println(substringXorQueries("1", [][]int{{4, 5}}))              //[[0 0]
}
