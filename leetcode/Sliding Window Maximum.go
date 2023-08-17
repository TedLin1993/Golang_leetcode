package leetcode

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, 0, n-k+1)
	deque := make([]int, 0, n)
	for i, num := range nums {
		if len(deque) > 0 && i-deque[0] >= k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= num {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			idx := deque[0]
			res = append(res, nums[idx])
		}
	}
	return res
}

func Test_maxSlidingWindow() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) //3,3,5,5,6,7
	fmt.Println(maxSlidingWindow([]int{1}, 1))                        //1
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))                  //7,4
}
