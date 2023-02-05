package leetcode

func isAlienSorted(words []string, order string) bool {
	orderMap := make([]int, 'z'-'a'+1)
	for i := 0; i < len(order); i++ {
		charIdx := order[i] - 'a'
		orderMap[charIdx] = i
	}
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}
			orderChar1, orderChar2 := orderMap[words[i][j]-'a'], orderMap[words[i+1][j]-'a']
			if orderChar1 < orderChar2 {
				break
			} else if orderChar1 > orderChar2 {
				return false
			}
		}
	}
	return true
}
