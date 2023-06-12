package leetcode

import "fmt"

func maxValue(n int, index int, maxSum int) int {
	if n == 1 {
		return maxSum
	}
	left, right := 1, maxSum
	for left < right {
		mid := left + (right-left)/2
		leftSum := 0
		if mid-index >= 1 {
			zeroVal := mid - index
			leftSum = (zeroVal+mid)*(index+1)/2 - mid
		} else {
			oneSum := index - mid + 1
			leftSum = oneSum + (mid+1)*mid/2 - mid
		}

		rightSum := 0
		if mid-(n-index-1) >= 1 {
			rightVal := mid - (n - index - 1)
			rightSum = (mid+rightVal)*(n-index)/2 - mid
		} else {
			oneSum := n - index - mid
			rightSum = (mid+1)*mid/2 + oneSum - mid
		}

		sum := leftSum + mid + rightSum
		if sum == maxSum {
			return mid
		} else if sum < maxSum {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

func Test_maxValue() {
	fmt.Println(maxValue(3, 0, 18)) //7
	fmt.Println(maxValue(3, 2, 18)) //7
	fmt.Println(maxValue(1, 0, 24)) //24
}
