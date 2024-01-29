package leetcode

func flowerGame(n int, m int) int64 {
	xOdd, xEven := n/2, n/2
	if n%2 == 1 {
		xOdd++
	}
	yOdd, yEven := m/2, m/2
	if m%2 == 1 {
		yOdd++
	}
	return int64(xOdd*yEven + xEven*yOdd)
}
