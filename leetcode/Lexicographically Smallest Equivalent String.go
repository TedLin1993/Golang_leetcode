package leetcode

import "fmt"

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	union := make([]int, 'z'-'a'+1)
	for idx := range union {
		union[idx] = idx
	}
	for i := 0; i < len(s1); i++ {
		sES_union(int(s1[i]-'a'), int(s2[i]-'a'), union)
	}
	res := make([]byte, len(baseStr))
	for idx := range baseStr {
		charIdx := sES_find(int(baseStr[idx]-'a'), union)
		res[idx] = byte(charIdx + 'a')
	}
	return string(res)
}

func sES_union(a, b int, union []int) {
	a, b = sES_find(a, union), sES_find(b, union)
	if a < b {
		union[b] = a
	} else {
		union[a] = b
	}
}

func sES_find(charIdx int, union []int) int {
	if union[charIdx] == charIdx {
		return charIdx
	}
	return sES_find(union[charIdx], union)
}

func Test_smallestEquivalentString() {
	fmt.Println(smallestEquivalentString("parker", "morris", "parser"))             //makkek
	fmt.Println(smallestEquivalentString("aaaaaaaaaa", "aaaaaaaaaa", "mzpxaabrxo")) //mzpxaabrxo
}
