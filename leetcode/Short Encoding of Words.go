package leetcode

import (
	"fmt"
	"sort"
)

type WordTrie struct {
	children [26]*WordTrie
}

func minimumLengthEncoding(words []string) int {
	root := WordTrie{}
	sort.Slice(words, func(a, b int) bool {
		return len(words[a]) > len(words[b])
	})

	res := 0
	for _, word := range words {
		isDistinct := false
		node := &root
		for i := len(word) - 1; i >= 0; i-- {
			idx := word[i] - 'a'
			if node.children[idx] == nil {
				isDistinct = true
				node.children[idx] = &WordTrie{}
			}
			node = node.children[idx]
		}
		if isDistinct {
			res += len(word) + 1
		}
	}
	return res
}

func TestminimumLengthEncoding() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"})) //10
	fmt.Println(minimumLengthEncoding([]string{"t"}))                  //2
}
