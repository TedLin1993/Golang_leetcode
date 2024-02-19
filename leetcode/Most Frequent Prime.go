package leetcode

import (
	"fmt"
)

func mostFrequentPrime(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	primeCount := make(map[int]int)
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, dir := range dirs {
				temp := 0
				cur_i, cur_j := i, j
				for cur_i >= 0 && cur_i < m && cur_j >= 0 && cur_j < n {
					temp = temp*10 + mat[cur_i][cur_j]
					if temp > 10 && isPrime(temp) {
						primeCount[temp]++
					}
					cur_i += dir[0]
					cur_j += dir[1]
				}
			}
		}
	}

	res := -1
	maxFreq := 0
	for prime, freq := range primeCount {
		if freq > maxFreq || (freq == maxFreq && prime > res) {
			res = prime
			maxFreq = freq
		}
	}
	return res
}

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Test_mostFrequentPrime() {
	// fmt.Println(mostFrequentPrime([][]int{{9, 7, 8}, {4, 6, 5}, {2, 8, 6}})) //97
	fmt.Println(mostFrequentPrime([][]int{{1}})) //-1
}
