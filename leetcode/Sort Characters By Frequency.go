package leetcode

import (
	"fmt"
	"strings"
)

func frequencySort(s string) string {
	charFreqMap := make(map[byte]int, len(s))
	for i := range s {
		charFreqMap[s[i]]++
	}

	FreqCountMap := make([][]byte, len(s)+1)
	for char, count := range charFreqMap {
		FreqCountMap[count] = append(FreqCountMap[count], char)
	}

	sb := strings.Builder{}
	sb.Grow(len(s))
	for count := len(FreqCountMap) - 1; count >= 0; count-- {
		if len(FreqCountMap[count]) > 0 {
			for _, c := range FreqCountMap[count] {
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
