package leetcode

func totalFruit(fruits []int) int {
	if len(fruits) <= 2 {
		return len(fruits)
	}
	type1, type2 := -1, -1
	curCount, continueCount := 0, 0
	res := 0
	for _, fruitType := range fruits {
		if fruitType != type1 && fruitType != type2 {
			res = max(res, curCount)
			type1, type2 = type2, fruitType
			curCount = continueCount + 1
			continueCount = 1
		} else if fruitType == type1 {
			type1, type2 = type2, type1
			curCount++
			continueCount = 1
		} else {
			curCount++
			continueCount++
		}
	}
	res = max(res, curCount)
	return res
}
