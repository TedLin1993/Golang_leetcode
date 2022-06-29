package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func minDeletions(s string) int {
	charCountList := make([]int, 26)
	for _, c := range s {
		charCountList[c-'a']++
	}
	sort.Slice(charCountList, func(i, j int) bool { return charCountList[i] > charCountList[j] })

	res := 0
	maxFreqAccepted := math.MaxInt16
	for i := range charCountList {
		if charCountList[i] == 0 {
			break
		}
		if charCountList[i] > maxFreqAccepted {
			res += charCountList[i] - maxFreqAccepted
			charCountList[i] = maxFreqAccepted
		}
		maxFreqAccepted = max(charCountList[i]-1, 0)
	}

	return res
}

func TestminDeletions() {
	fmt.Println(minDeletions("aab"))      //0
	fmt.Println(minDeletions("aaabbbcc")) //2
	fmt.Println(minDeletions("ceabaacb")) //2
	fmt.Println(minDeletions("bbcebab"))  //2
}
