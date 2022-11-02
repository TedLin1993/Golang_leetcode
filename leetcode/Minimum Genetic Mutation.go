package leetcode

import "fmt"

func minMutation(start string, end string, bank []string) int {
	if end == start {
		return 0
	}
	isVisited := make([]bool, len(bank))
	queue := make([]string, 0, len(bank))
	queue = append(queue, start)
	res := 0
	for len(queue) > 0 {
		res++
		length := len(queue)
		for i := 0; i < length; i++ {
			temp := queue[i]
			for j := 0; j < len(bank); j++ {
				if !isVisited[j] && canMutate(temp, bank[j]) {
					if end == bank[j] {
						return res
					}
					queue = append(queue, bank[j])
					isVisited[j] = true
				}
			}
		}
		queue = queue[length:]
	}
	return -1
}

func canMutate(start, next string) bool {
	count := 0
	for i := 0; i < len(start); i++ {
		if start[i] != next[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return count == 1
}

func Test_minMutation() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))                         //1
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"})) //2
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"})) //3
}
