package leetcode

func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow(1/x, -n)
	}
	res := 1.0
	for n > 0 {
		if n%2 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}
