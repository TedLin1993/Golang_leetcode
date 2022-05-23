package leetcode

import (
	"fmt"
	"strings"
)

func minRemoveToMakeValid(s string) string {
	stringList := []byte(s)
	leftBucketStack := []int{}
	for i := 0; i < len(stringList); i++ {
		if stringList[i] == '(' {
			leftBucketStack = append(leftBucketStack, i)
		} else if stringList[i] == ')' {
			if len(leftBucketStack) > 0 {
				leftBucketStack = leftBucketStack[:len(leftBucketStack)-1]
			} else {
				stringList[i] = '#'
			}
		}
	}
	for len(leftBucketStack) > 0 {
		index := leftBucketStack[len(leftBucketStack)-1]
		stringList[index] = '#'
		leftBucketStack = leftBucketStack[:len(leftBucketStack)-1]
	}

	return strings.ReplaceAll(string(stringList), "#", "")
}

func TestminRemoveToMakeValid() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)")) //"lee(t(c)o)de"
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))       //"ab(c)d"
	fmt.Println(minRemoveToMakeValid("))(("))          //""
}
