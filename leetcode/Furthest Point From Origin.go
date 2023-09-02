package leetcode

func furthestDistanceFromOrigin(moves string) int {
	n := len(moves)
	lCount, rCount := 0, 0
	for i := 0; i < n; i++ {
		if moves[i] == 'L' {
			lCount++
		} else if moves[i] == 'R' {
			rCount++
		}
	}
	sCount := n - lCount - rCount
	if lCount > rCount {
		return lCount - rCount + sCount
	}
	return rCount - lCount + sCount
}
