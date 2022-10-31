package leetcode

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[[26]int][]string, len(strs))
	for _, str := range strs {
		charArr := [26]int{}
		for i := 0; i < len(str); i++ {
			charArr[str[i]-'a']++
		}
		strMap[charArr] = append(strMap[charArr], str)
	}

	// return maps.Values(strMap)

	res := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		res = append(res, v)
	}
	return res
}

func Test_groupAnagrams() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) //[[eat tea ate] [tan nat] [bat]]
	fmt.Println(groupAnagrams([]string{""}))                                       //[[]]
	fmt.Println(groupAnagrams([]string{"a"}))                                      //[[a]]
}
