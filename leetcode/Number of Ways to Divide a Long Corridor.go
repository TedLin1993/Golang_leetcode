package leetcode

func numberOfWays(corridor string) int {
	sCount := 0
	for _, c := range corridor {
		if c == 'S' {
			sCount++
		}
	}
	if sCount < 2 || sCount%2 == 1 {
		return 0
	}
	if sCount == 2 {
		return 1
	}
	res := 1
	mod := int(1e9 + 7)
	dividedCount := 1
	curS := 0
	for _, c := range corridor {
		if c == 'S' {
			if curS < 2 {
				curS++
			} else {
				res = res * dividedCount % mod
				curS = 1
				dividedCount = 1
			}
		} else {
			if curS < 2 {
				continue
			}
			dividedCount++
		}
	}
	return res
}
