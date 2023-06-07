package leetcode

func minFlips(a int, b int, c int) int {
	res := 0
	for a != 0 || b != 0 || c != 0 {
		ra, rb, rc := a%2, b%2, c%2
		if rc == 0 {
			res += ra + rb
		} else {
			if ra == 0 && rb == 0 {
				res++
			}
		}
		a, b, c = a>>1, b>>1, c>>1
	}
	return res
}
