package leetcode

import "fmt"

var partitionMap map[string][][]string

func partitionPalinrome(s string) [][]string {
	partitionMap = make(map[string][][]string)
	return partitionRecur(s)
}

func partitionRecur(s string) [][]string {
	res := [][]string{}
	for i := 1; i <= len(s); i++ {
		if !isPalindrome(s[0:i]) {
			continue
		}
		if i == len(s) {
			res = append(res, []string{s[0:i]})
			partitionMap[s] = res
			return res
		}
		if rightPartition, ok := partitionMap[s[i:]]; ok {
			for j := 0; j < len(rightPartition); j++ {
				temp := append([]string{s[0:i]}, rightPartition[j]...)
				res = append(res, temp)
			}
		} else {
			rightPartition := partitionPalinrome(s[i:])
			for j := 0; j < len(rightPartition); j++ {
				temp := append([]string{s[0:i]}, rightPartition[j]...)
				res = append(res, temp)
			}
		}
	}
	partitionMap[s] = res
	return res
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func TestPalindrome_Partitioning() {
	fmt.Println(partitionPalinrome("aab"))  //[[a a b] [aa b]]
	fmt.Println(partitionPalinrome("a"))    //[[a]]
	fmt.Println(partitionPalinrome("abba")) //[[a b b a] [a bb a] [abba]]
	fmt.Println(partitionPalinrome("bbab")) //[[b b a b] [b bab] [bb a b]]
	fmt.Println(partitionPalinrome("bbb"))  //[[b b b] [b bb] [bb b] [bbb]]
}
