package leetcode

import (
	"fmt"
)

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res, _ := findTiltRecur(root)
	return res
}

func findTiltRecur(node *TreeNode) (resultSum int, originSum int) {
	if node == nil {
		return 0, 0
	}
	resLeftSum, originLeftSum := findTiltRecur(node.Left)
	resRightSum, originRightSum := findTiltRecur(node.Right)
	originSum = node.Val + originLeftSum + originRightSum
	node.Val = abs(originLeftSum - originRightSum)
	resultSum = node.Val + resLeftSum + resRightSum
	return
}

func TestfindTilt() {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}
	node1.Left = &node2
	node2.Left = &node3
	node3.Left = &node4
	node4.Left = &node5
	fmt.Println(findTilt(&node1))
}
