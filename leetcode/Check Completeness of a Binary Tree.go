package leetcode

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	isNil := false
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			nodes := []*TreeNode{queue[i].Left, queue[i].Right}
			for _, node := range nodes {
				if !isNil && node != nil {
					queue = append(queue, node)
				} else if !isNil && node == nil {
					isNil = true
				} else if isNil && node != nil {
					return false
				}
			}
		}
		queue = queue[length:]
	}
	return true
}
