package leetcode

import "fmt"

func widthOfBinaryTree(root *TreeNode) int {
	root.Val = 0
	queue := []*TreeNode{root}
	res := 1
	for len(queue) > 0 {
		layerNodeCount := len(queue)
		leftIdx, rightIdx := -1, -1
		for i := 0; i < layerNodeCount; i++ {
			if queue[i] != nil {
				if leftIdx == -1 {
					leftIdx = queue[i].Val
					rightIdx = queue[i].Val
				} else {
					rightIdx = queue[i].Val
				}
				if queue[i].Left != nil {
					queue[i].Left.Val = queue[i].Val << 1
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue[i].Right.Val = queue[i].Val<<1 + 1
					queue = append(queue, queue[i].Right)
				}
			}
		}
		width := rightIdx - leftIdx + 1
		if width > res {
			res = width
		}
		queue = queue[layerNodeCount:]
	}
	return res
}

func TestWidthOfBinaryTree() {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}
	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node3.Right = &node5
	fmt.Println(widthOfBinaryTree(&node1))
}
