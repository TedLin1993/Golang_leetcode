package leetcode

import "fmt"

func vowelStrings(words []string, queries [][]int) []int {
	prefixSum := make([]int, len(words))
	if isVowel(words[0]) {
		prefixSum[0] = 1
	}
	for i := 1; i < len(words); i++ {
		if isVowel(words[i]) {
			prefixSum[i] = prefixSum[i-1] + 1
		} else {
			prefixSum[i] = prefixSum[i-1]
		}
	}

	res := make([]int, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		if start == 0 {
			res[i] = prefixSum[end]
		} else {
			res[i] = prefixSum[end] - prefixSum[start-1]
		}
	}
	return res
}

func isVowel(word string) bool {
	vowel := []byte{'a', 'e', 'i', 'o', 'u'}
	isStartVowel, isEndVowel := false, false
	for i := 0; i < len(vowel); i++ {
		if word[0] == vowel[i] {
			isStartVowel = true
		}
		if word[len(word)-1] == vowel[i] {
			isEndVowel = true
		}
	}

	return isStartVowel && isEndVowel
}

func Test_vowelStrings() {
	fmt.Println(vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
}
