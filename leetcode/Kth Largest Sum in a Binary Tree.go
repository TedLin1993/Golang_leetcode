package leetcode

import "sort"

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levelSums := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		sum := 0
		for i := 0; i < length; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		levelSums = append(levelSums, sum)
		queue = queue[length:]
	}
	if k > len(levelSums) {
		return -1
	}
	sort.Slice(levelSums, func(i, j int) bool {
		return levelSums[i] > levelSums[j]
	})
	return int64(levelSums[k-1])
}
