package leetcode

import "fmt"

func firstUniqChar(s string) int {
	charMap := make([]int, 26)
	for i := 0; i < len(s); i++ {
		charMap[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if charMap[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func Test_firstUniqChar() {
	fmt.Println(firstUniqChar("leetcode"))     //0
	fmt.Println(firstUniqChar("loveleetcode")) //2
	fmt.Println(firstUniqChar("aabb"))         //-1
}
