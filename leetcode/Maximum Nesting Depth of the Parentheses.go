package leetcode

func maxDepth_(s string) int {
	res := 0
	cur := 0
	for _, c := range s {
		if c == '(' {
			cur++
			res = max(res, cur)
		} else if c == ')' {
			cur--
		}
	}
	return res
}
