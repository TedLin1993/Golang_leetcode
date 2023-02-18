package leetcode

import (
	"fmt"
)

func substringXorQueries(s string, queries [][]int) [][]int {
	res := make([][]int, 0, len(queries))
	valueMap := make(map[int][]int)
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			if _, ok := valueMap[0]; !ok {
				valueMap[0] = []int{i, i}
			}
			continue
		}
		value := 0
		for j := i; j < min(i+30, len(s)); j++ {
			value = value<<1 + int(s[j]-'0')
			if _, ok := valueMap[value]; !ok {
				valueMap[value] = []int{i, j}
			}
		}
	}
	for _, query := range queries {
		value := query[0] ^ query[1]
		v, ok := valueMap[value]
		if !ok {
			res = append(res, []int{-1, -1})
		} else {
			res = append(res, v)
		}
	}
	return res

}

func Test_substringXorQueries() {
	fmt.Println(substringXorQueries("101101", [][]int{{0, 5}, {1, 2}})) //[[0 2] [2 3]]
	fmt.Println(substringXorQueries("0101", [][]int{{12, 8}}))          //[[-1 -1]
	fmt.Println(substringXorQueries("1", [][]int{{4, 5}}))              //[[0 0]
}
