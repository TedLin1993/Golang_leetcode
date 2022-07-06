package leetcode

import "fmt"

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	up, down, peak := 0, 0, 0
	res := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			up++
			peak = up + 1
			down = 0
			res += up + 1
		} else if ratings[i] < ratings[i-1] {
			down++
			up = 0
			if down >= peak {
				res += down + 1
			} else {
				res += down
			}
		} else {
			up = 0
			down = 0
			peak = 0
			res += 1
		}
	}
	return res
}

func Testcandy() {
	fmt.Println(candy([]int{1, 0, 2}))       //5
	fmt.Println(candy([]int{1, 2, 2}))       //4
	fmt.Println(candy([]int{1, 2, 3, 2, 1})) //9
}
