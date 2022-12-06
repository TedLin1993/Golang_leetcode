package leetcode

import (
	"fmt"
	"strings"
)

func frequencySort(s string) string {
	freqMap := make(map[byte]int, len(s))
	for i := range s {
		freqMap[s[i]]++
	}

	countCharMap := make([][]byte, len(s)+1)
	for char, count := range freqMap {
		countCharMap[count] = append(countCharMap[count], char)
	}

	sb := strings.Builder{}
	sb.Grow(len(s))
	for count := len(countCharMap) - 1; count >= 0; count-- {
		if len(countCharMap[count]) > 0 {
			for _, c := range countCharMap[count] {
				for i := 0; i < count; i++ {
					sb.WriteByte(c)
				}
			}
		}
	}
	return sb.String()
}

func Test_frequencySort() {
	fmt.Println(frequencySort("tree"))                        //eert
	fmt.Println(frequencySort("cccaaa"))                      //aaaccc
	fmt.Println(frequencySort("Aabb"))                        //bbAa
	fmt.Println(frequencySort("2a554442f544asfasssffffasss")) //sssssssffffff44444aaaa55522
}
