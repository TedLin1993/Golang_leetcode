package leetcode

import (
	"fmt"
	"strconv"
)

func tree2str(node *TreeNode) string {
	if node == nil {
		return "()"
	}

	valStr := strconv.Itoa(node.Val)
	leftStr := tree2str(node.Left)
	rightStr := tree2str(node.Right)
	if leftStr == "()" && rightStr == "()" {
		return valStr
	}
	if rightStr == "()" {
		return fmt.Sprintf("%s(%s)", valStr, leftStr)
	}
	if leftStr == "()" {
		return fmt.Sprintf("%s()(%s)", valStr, rightStr)
	}
	return fmt.Sprintf("%s(%s)(%s)", valStr, leftStr, rightStr)
}
