package leetcode

import "fmt"

func numMatchingSubseq(s string, words []string) int {
	type entry struct {
		WordsIdx int
		CharIdx  int
	}

	//prefix array
	nextEntries := make([][]entry, 26)
	for wordsIdx, word := range words {
		charIdx := word[0] - 'a'
		nextEntries[charIdx] = append(nextEntries[charIdx], entry{WordsIdx: wordsIdx, CharIdx: 0})
	}

	res := 0
	for i := 0; i < len(s); i++ {
		charIdx := s[i] - 'a'
		entry := nextEntries[charIdx]
		nextEntries[charIdx] = nil
		for j := 0; j < len(entry); j++ {
			entry[j].CharIdx++
			wordsIdx := entry[j].WordsIdx
			charIdx := entry[j].CharIdx
			if entry[j].CharIdx == len(words[wordsIdx]) {
				res++
				continue
			}
			nextCharIdx := words[wordsIdx][charIdx] - 'a'
			nextEntries[nextCharIdx] = append(nextEntries[nextCharIdx], entry[j])
		}
	}

	return res
}

func TestnumMatchingSubseq() {
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))                          //3
	fmt.Println(numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"})) //2
}
