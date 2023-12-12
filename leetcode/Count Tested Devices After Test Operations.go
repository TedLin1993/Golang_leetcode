package leetcode

func countTestedDevices(batteryPercentages []int) int {
	testCount := 0
	for _, b := range batteryPercentages {
		if b-testCount > 0 {
			testCount++
		}
	}
	return testCount
}
