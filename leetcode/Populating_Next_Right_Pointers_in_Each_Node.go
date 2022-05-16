package leetcode

import "fmt"

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	levelFirstNode := root
	for levelFirstNode.Left != nil {
		currentNode := levelFirstNode
		for currentNode != nil {
			currentNode.Left.Next = currentNode.Right
			if currentNode.Next != nil {
				currentNode.Right.Next = currentNode.Next.Left
			}
			currentNode = currentNode.Next
		}
		levelFirstNode = levelFirstNode.Left
	}

	return root
}

func TestPopulating_Next_Right_Pointers_in_Each_Node() {
	node1 := Node{Val: 1}
	node2 := Node{Val: 2}
	node3 := Node{Val: 3}
	node4 := Node{Val: 4}
	node5 := Node{Val: 5}
	node6 := Node{Val: 6}
	node7 := Node{Val: 7}
	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node2.Right = &node5
	node3.Left = &node6
	node3.Right = &node7
	fmt.Println(connect(&node1))
}
