package leetcode

type pair struct{ pre, suf byte }
type trie struct {
	next  map[pair]*trie
	count int
}

func countPrefixSuffixPairs_3045(words []string) int64 {
	res := 0
	root := &trie{next: map[pair]*trie{}}
	for _, w := range words {
		cur := root
		for i := range w {
			p := pair{w[i], w[len(w)-1-i]}
			if cur.next[p] == nil {
				cur.next[p] = &trie{next: map[pair]*trie{}}
			}
			cur = cur.next[p]
			res += cur.count
		}
		cur.count++
	}
	return int64(res)
}
