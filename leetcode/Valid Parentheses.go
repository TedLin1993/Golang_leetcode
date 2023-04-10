package leetcode

import "fmt"

func isValid(s string) bool {
	bucketStack := []byte{}
	for _, bucket := range s {
		if bucket == '(' || bucket == '{' || bucket == '[' {
			bucketStack = append(bucketStack, byte(bucket))
		} else {
			if len(bucketStack) == 0 || !checkBucket(byte(bucket), bucketStack[len(bucketStack)-1]) {
				return false
			} else {
				bucketStack = bucketStack[:len(bucketStack)-1]
			}
		}
	}
	return len(bucketStack) == 0
}

func checkBucket(rightBucket byte, leftBucket byte) bool {
	if (rightBucket == ')' && leftBucket == '(') ||
		(rightBucket == '}' && leftBucket == '{') ||
		(rightBucket == ']' && leftBucket == '[') {
		return true
	}
	return false
}

func isValid2(s string) bool {
	stack := make([]byte, 0, len(s)/2)
	bracketMap := map[byte]byte{')': '(', '}': '{', ']': '['}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 || bracketMap[s[i]] != stack[len(stack)-1] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func TestisValid() {
	fmt.Println(isValid("()"))     //true
	fmt.Println(isValid("()[]{}")) //true
	fmt.Println(isValid("(]"))     //false
}
