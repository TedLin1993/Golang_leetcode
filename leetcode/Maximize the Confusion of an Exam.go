package leetcode

func maxConsecutiveAnswers(answerKey string, k int) int {
	left := 0
	res := 0
	tCount, fCount := 0, 0
	for right := 0; right < len(answerKey); right++ {
		if answerKey[right] == 'T' {
			tCount++
		} else {
			fCount++
		}
		for min(tCount, fCount) > k {
			if answerKey[left] == 'T' {
				tCount--
			} else {
				fCount--
			}
			left++
		}
		res = max(res, right-left+1)
	}
	res = max(res, len(answerKey)-left)
	return res
}
