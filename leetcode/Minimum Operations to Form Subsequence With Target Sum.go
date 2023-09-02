package leetcode

func minOperations_2835(nums []int, target int) int {
	sum := 0
	powerMap := make([]int, 32)
	for _, num := range nums {
		sum += num
		power := getPower(num)
		powerMap[power]++
	}
	if sum < target {
		return -1
	}
	res := 0
	curPower := 0
	for target > 0 {
		remain := target % 2
		if remain == 1 {
			if powerMap[curPower] == 0 {
				powerMap[curPower] += 2
				for i := curPower + 1; i <= 30; i++ {
					res++
					if powerMap[i] > 0 {
						powerMap[i]--
						break
					} else {
						powerMap[i]++
					}
				}
			}
			powerMap[curPower]--
		}
		powerMap[curPower+1] += powerMap[curPower] / 2
		target = target >> 1
		curPower++
	}
	return res
}

func getPower(num int) int {
	power := 0
	num = num >> 1
	for num != 0 {
		num = num >> 1
		power++
	}
	return power
}
