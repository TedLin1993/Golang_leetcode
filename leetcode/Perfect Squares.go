package leetcode

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	sqrt_n := int(math.Sqrt(float64(n)))
	squrtNums := make([]int, 0, sqrt_n)
	for i := sqrt_n; i > 0; i-- {
		squrtNums = append(squrtNums, i*i)
	}

	visit := make([]bool, n)
	queue := []int{n}
	res := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i] == 0 {
				return res
			}
			for j := 0; j < len(squrtNums); j++ {
				num := queue[i] - squrtNums[j]
				if num < 0 || visit[num] {
					continue
				}
				visit[num] = true
				queue = append(queue, num)
			}
		}
		res++
		queue = queue[length:]
	}
	return res
}

func Test_numSquares() {
	fmt.Println(numSquares(3))     //3
	fmt.Println(numSquares(12))    //3
	fmt.Println(numSquares(13))    //2
	fmt.Println(numSquares(48))    //3
	fmt.Println(numSquares(50))    //2
	fmt.Println(numSquares(10000)) //1
	fmt.Println(numSquares(9999))  //4
}
