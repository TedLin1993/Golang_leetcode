package leetcode

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	//find middle node
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	midNode := slow

	//reverse listnode
	lastNode := midNode
	node := midNode
	nextNode := node.Next
	node.Next = nil
	for nextNode != nil {
		lastNode = nextNode
		node, nextNode, nextNode.Next = nextNode, nextNode.Next, node
	}

	//check palindrome
	left := head
	right := lastNode
	for left != midNode {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
