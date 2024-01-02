package leetcode

import (
	"fmt"
)

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		node := &TreeNode{Val: val, Left: root}
		return node
	}

	queue := []*TreeNode{root}
	for i := 2; i < depth; i++ {
		length := len(queue)
		for j := 0; j < length; j++ {
			if queue[j].Left != nil {
				queue = append(queue, queue[j].Left)
			}
			if queue[j].Right != nil {
				queue = append(queue, queue[j].Right)
			}
		}
		queue = queue[length:]
	}

	for _, node := range queue {
		left := &TreeNode{Val: val, Left: node.Left}
		right := &TreeNode{Val: val, Right: node.Right}
		node.Left = left
		node.Right = right
	}
	return root
}

func Test_addOneRow() {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	fmt.Println(addOneRow(&node1, 5, 4))
}
