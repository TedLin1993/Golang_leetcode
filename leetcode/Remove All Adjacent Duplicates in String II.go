package leetcode

import "fmt"

func removeDuplicatesII(s string, k int) string {
	stack := []string{string(s[0])}
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 {
			if s[i] == stack[len(stack)-1][0] {
				stack[len(stack)-1] += string(s[i])
				if len(stack[len(stack)-1]) == k {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, string(s[i]))
			}
		} else {
			stack = append(stack, string(s[i]))
		}

	}

	var res string
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}

func TestremoveDuplicatesII() {
	fmt.Println(removeDuplicatesII("aaa", 3))                  //""
	fmt.Println(removeDuplicatesII("abcd", 2))                 //"abcd"
	fmt.Println(removeDuplicatesII("deeedbbcccbdaa", 3))       //"aa"
	fmt.Println(removeDuplicatesII("pbbcggttciiippooaais", 2)) //"ps"
}
