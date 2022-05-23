package leetcode

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Neighbors == nil || len(node.Neighbors) == 0 {
		return &Node{Val: node.Val}
	}

	nodeMap := map[int]*Node{}
	queue := []*Node{node}
	for len(queue) > 0 {
		originNode := queue[0]
		queue = queue[1:]
		if _, ok := nodeMap[originNode.Val]; ok {
			continue
		}

		cloneNode := &Node{Val: originNode.Val}
		nodeMap[cloneNode.Val] = cloneNode

		for _, neighbor := range originNode.Neighbors {
			if _, ok := nodeMap[neighbor.Val]; !ok {
				queue = append(queue, neighbor)
				continue
			}
			cloneNode.Neighbors = append(cloneNode.Neighbors, nodeMap[neighbor.Val])
			nodeMap[neighbor.Val].Neighbors = append(nodeMap[neighbor.Val].Neighbors, cloneNode)
		}
	}

	return nodeMap[node.Val]
}
