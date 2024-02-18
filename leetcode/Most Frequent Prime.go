package leetcode

import (
	"fmt"
	"math"
)

func mostFrequentPrime(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	isPrimeList := getPrimeList(int(math.Pow10(max(m, n))))
	primeCount := map[int]int{}
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, dir := range dirs {
				temp := 0
				cur_i, cur_j := i, j
				for cur_i >= 0 && cur_i < m && cur_j >= 0 && cur_j < n {
					temp = temp*10 + mat[cur_i][cur_j]
					if temp > 10 && isPrimeList[temp] {
						primeCount[temp]++
					}
					cur_i += dir[0]
					cur_j += dir[1]
				}
			}
		}
	}

	if len(primeCount) == 0 {
		return -1
	}
	res := 0
	maxFreq := 0
	for k, v := range primeCount {
		if v > maxFreq {
			res = k
			maxFreq = v
		} else if v == maxFreq {
			res = max(res, k)
		}
	}
	return res
}

// func getPrimeList(n int) []bool {
// 	list := make([]bool, n+1)
// 	for i := 2; i < len(list); i++ {
// 		list[i] = true
// 	}
// 	for i := 2; i < len(list); i++ {
// 		for j := i * i; j < len(list); j += i {
// 			list[j] = false
// 		}
// 	}
// 	return list
// }

func Test_mostFrequentPrime() {
	// fmt.Println(mostFrequentPrime([][]int{{9, 7, 8}, {4, 6, 5}, {2, 8, 6}})) //97
	fmt.Println(mostFrequentPrime([][]int{{1}})) //-1
}
