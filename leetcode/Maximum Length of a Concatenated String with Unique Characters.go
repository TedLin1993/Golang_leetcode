package leetcode

import "fmt"

func maxLength(arr []string) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res = max(res, dfs_maxLength(arr, i, [26]bool{}))
	}
	return res
}

func dfs_maxLength(arr []string, index int, uniqChars [26]bool) int {
	str := arr[index]
	for j := 0; j < len(str); j++ {
		charIdx := str[j] - 'a'
		if uniqChars[charIdx] {
			return 0
		}
		uniqChars[charIdx] = true
	}

	res := 0
	for i := index + 1; i < len(arr); i++ {
		res = max(res, dfs_maxLength(arr, i, uniqChars))
	}
	return res + len(str)
}

func Test_maxLength() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))                                                               //4
	fmt.Println(maxLength([]string{"r", "cha", "act", "ers"}))                                                       //6
	fmt.Println(maxLength([]string{"aa", "bb"}))                                                                     //0
	fmt.Println(maxLength([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"})) //16
}
