package leetcode

import "fmt"

var strMap map[string]string

func numSimilarGroups(strs []string) int {
	strMap = make(map[string]string, len(strs))
	for _, str := range strs {
		strMap[str] = str
	}
	for i := 0; i < len(strs)-1; i++ {
		for j := i + 1; j < len(strs); j++ {
			union_numSimilarGroups(strs[i], strs[j])
		}
	}
	for i := 0; i < len(strs); i++ {
		find_numSimilarGroups(strs[i])
	}
	checkMap := map[string]bool{}
	for _, str := range strMap {
		if !checkMap[str] {
			checkMap[str] = true
		}
	}
	return len(checkMap)
}

func find_numSimilarGroups(str string) string {
	if strMap[str] == str {
		return str
	}
	strMap[str] = find_numSimilarGroups(strMap[str])
	return strMap[str]
}

func union_numSimilarGroups(str1, str2 string) {
	if find_numSimilarGroups(str1) == find_numSimilarGroups(str2) {
		return
	}
	misMatchCount := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			misMatchCount++
		}
	}
	if misMatchCount == 2 {
		strMap[find_numSimilarGroups(str2)] = find_numSimilarGroups(str1)
	}
}

func Test_numSimilarGroups() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"}))                                                                                                             //2
	fmt.Println(numSimilarGroups([]string{"omv", "ovm"}))                                                                                                                               //1
	fmt.Println(numSimilarGroups([]string{"kccomwcgcs", "socgcmcwkc", "sgckwcmcoc", "coswcmcgkc", "cowkccmsgc", "cosgmccwkc", "sgmkwcccoc", "coswmccgkc", "kowcccmsgc", "kgcomwcccs"})) //5
}
