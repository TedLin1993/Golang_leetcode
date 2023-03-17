package leetcode

type TrieNode struct {
	nextChars [26]*TrieNode
	isEnd     bool
}

func Constructor_TrieNode() TrieNode {
	return TrieNode{}
}

func (this *TrieNode) Insert(word string) {
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if this.nextChars[idx] == nil {
			this.nextChars[idx] = &TrieNode{}
		}
		this = this.nextChars[idx]
	}
	this.isEnd = true
}

func (this *TrieNode) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if this.nextChars[idx] == nil {
			return false
		}
		this = this.nextChars[idx]
	}
	return this.isEnd
}

func (this *TrieNode) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if this.nextChars[idx] == nil {
			return false
		}
		this = this.nextChars[idx]
	}
	return true
}
