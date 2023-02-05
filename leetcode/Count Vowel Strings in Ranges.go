package leetcode

import "fmt"

func vowelStrings(words []string, queries [][]int) []int {
	isVowels := make([]bool, len(words))
	for i, word := range words {
		if isVowel(word) {
			isVowels[i] = true
		}
	}

	res := make([]int, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		for j := start; j <= end; j++ {
			if isVowels[j] {
				res[i]++
			}
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
