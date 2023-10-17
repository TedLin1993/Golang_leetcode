package leetcode

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	var find func(idx int) int
	find = func(idx int) int {
		if parents[idx] == idx {
			return idx
		}
		parents[idx] = find(parents[idx])
		return parents[idx]
	}
	union := func(parent, child int) {
		parent, child = find(parent), find(child)
		parents[child] = parent
	}
	for i := 0; i < n; i++ {
		left, right := leftChild[i], rightChild[i]
		if left != -1 {
			if find(left) != left || find(i) == left {
				return false
			}
			union(i, left)
		}
		if right != -1 {
			if find(right) != right || find(i) == right {
				return false
			}
			union(i, right)
		}
	}
	for i := 1; i < n; i++ {
		if find(i) != find(i-1) {
			return false
		}
	}
	return true
}
