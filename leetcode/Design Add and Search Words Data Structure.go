package leetcode

type Trie2 struct {
	children [26]*Trie2
	isEnd    bool
}

type WordDictionary struct {
	trie *Trie2
}

func Constructor_211() WordDictionary {
	return WordDictionary{trie: &Trie2{}}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.trie
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &Trie2{}
		}
		cur = cur.children[idx]
	}
	cur.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(word string, idx int, cur *Trie2) bool
	dfs = func(word string, idx int, cur *Trie2) bool {
		if cur == nil {
			return false
		}
		if idx == len(word) {
			return cur.isEnd
		}
		if word[idx] == '.' {
			for i := 0; i < len(cur.children); i++ {
				if dfs(word, idx+1, cur.children[i]) {
					return true
				}
			}
			return false
		}
		return dfs(word, idx+1, cur.children[word[idx]-'a'])
	}
	return dfs(word, 0, this.trie)
}
