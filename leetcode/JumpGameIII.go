package leetcode

func canReach(arr []int, start int) bool {
	isReached := make([]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		isReached[i] = false
	}
	return canReachRecur(arr, start, &isReached)
}

func canReachRecur(arr []int, currentIndex int, isReached *[]bool) bool {
	if currentIndex < 0 || currentIndex >= len(arr) || (*isReached)[currentIndex] {
		return false
	}
	if arr[currentIndex] == 0 {
		return true
	}

	(*isReached)[currentIndex] = true
	return canReachRecur(arr, currentIndex-arr[currentIndex], isReached) || canReachRecur(arr, currentIndex+arr[currentIndex], isReached)
}

func TestJumpGameIII() {
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	start := 5
	println(canReach(arr, start))

	arr = []int{4, 2, 3, 0, 3, 1, 2}
	start = 0
	println(canReach(arr, start))

	arr = []int{3, 0, 2, 1, 2}
	start = 2
	println(canReach(arr, start))

	arr = []int{31, 70, 40, 4, 27, 28, 44, 17, 8, 16, 64, 14, 30, 17, 25, 61, 33, 50, 45, 67, 12, 43, 72, 0, 26, 24, 33, 66, 22, 11, 61, 15, 2, 18, 51, 37, 34, 7, 14, 29, 37, 3, 40, 3, 61, 20, 24, 22, 53, 18, 58, 56, 16, 44, 14, 5, 6, 19, 72, 46, 49, 10, 42, 40, 46, 10, 6, 34, 17, 68, 62, 34, 33, 68}
	start = 10
	println(canReach(arr, start))
}
