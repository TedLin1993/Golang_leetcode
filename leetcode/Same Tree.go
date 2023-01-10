package leetcode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	p_queue := []*TreeNode{p}
	q_queue := []*TreeNode{q}
	for len(p_queue) > 0 || len(q_queue) > 0 {
		p0 := p_queue[0]
		q0 := q_queue[0]
		if p0 == nil || q0 == nil {
			if p0 != q0 {
				return false
			}
		} else {
			if p0.Val != q0.Val {
				return false
			}
			p_queue = append(p_queue, p0.Left)
			p_queue = append(p_queue, p0.Right)
			q_queue = append(q_queue, q0.Left)
			q_queue = append(q_queue, q0.Right)
		}

		p_queue = p_queue[1:]
		q_queue = q_queue[1:]
	}
	return len(p_queue) == 0 && len(q_queue) == 0
}
