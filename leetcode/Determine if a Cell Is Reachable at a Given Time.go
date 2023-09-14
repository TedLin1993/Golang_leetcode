package leetcode

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	xDiff, yDiff := abs(fx-sx), abs(fy-sy)
	minStep := min(xDiff, yDiff)
	xDiff -= minStep
	yDiff -= minStep
	minStep += xDiff + yDiff
	if minStep == 0 && t == 1 {
		return false
	}
	return minStep <= t
}
