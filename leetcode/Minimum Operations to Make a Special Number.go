package leetcode

func minimumOperations(num string) int {
	n := len(num)
	if n <= 1 {
		return 1
	}
	lastTwo := []string{"00", "25", "50", "75"}
	res := n
	for _, d := range num {
		if d == '0' {
			res = n - 1
			break
		}
	}
	for _, lt := range lastTwo {
		n1, n2 := lt[0], lt[1]
		idx1 := -1
		for i := n - 1; i >= 0; i-- {
			if num[i] == n2 {
				idx1 = i
				break
			}
		}
		idx2 := -1
		for i := idx1 - 1; i >= 0; i-- {
			if num[i] == n1 {
				idx2 = i
				break
			}
		}
		if idx2 == -1 {
			continue
		}
		res = min(res, n-idx2-2)
	}
	return res
}
