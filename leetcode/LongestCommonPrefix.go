package leetcode

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	res := ""

	for i := 0; ; i++ {
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[j][i] != strs[0][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}
}

func Test_longestCommonPrefix() {
	var strs []string
	strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs)) //"fl"

	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs)) //""

	strs = []string{"abc", "abc", "abc"}
	fmt.Println(longestCommonPrefix(strs)) //""
}
