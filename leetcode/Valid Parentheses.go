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

func TestisValid() {
	fmt.Println(isValid("()"))     //true
	fmt.Println(isValid("()[]{}")) //true
	fmt.Println(isValid("(]"))     //false
}
