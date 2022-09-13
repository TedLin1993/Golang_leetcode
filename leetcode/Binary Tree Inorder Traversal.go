package leetcode

import "fmt"

func inorderTraversal(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	res := []int{}
	stack := []*TreeNode{node}
	for len(stack) > 0 {
		lastNode := stack[len(stack)-1]
		if lastNode.Left != nil {
			stack = append(stack, lastNode.Left)
			lastNode.Left = nil
			continue
		}
		res = append(res, lastNode.Val)
		stack = stack[:len(stack)-1]
		if lastNode.Right != nil {
			stack = append(stack, lastNode.Right)
		}
	}
	return res
}

func Test_inorderTraversal() {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node1.Right = &node2
	node2.Left = &node3
	fmt.Println(inorderTraversal(&node1))
}
