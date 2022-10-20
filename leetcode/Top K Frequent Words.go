package leetcode

import (
	"container/heap"
	"fmt"
)

type wordFreq struct {
	Word  string
	Count int
}

type MinWordFreqHeap []wordFreq

func (h MinWordFreqHeap) Len() int { return len(h) }
func (h MinWordFreqHeap) Less(i, j int) bool {
	if h[i].Count != h[j].Count {
		return h[i].Count < h[j].Count
	}
	return h[i].Word > h[j].Word
}
func (h MinWordFreqHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinWordFreqHeap) Push(x interface{}) {
	*h = append(*h, x.(wordFreq))
}
func (h *MinWordFreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	wordMap := make(map[string]int, len(words))
	for _, word := range words {
		wordMap[word]++
	}

	minwordFreqHeap := MinWordFreqHeap{}
	for word, count := range wordMap {
		if len(minwordFreqHeap) < k {
			heap.Push(&minwordFreqHeap, wordFreq{Word: word, Count: count})
		} else {
			if count < minwordFreqHeap[0].Count ||
				(count == minwordFreqHeap[0].Count && word > minwordFreqHeap[0].Word) {
				continue
			}
			heap.Pop(&minwordFreqHeap)
			heap.Push(&minwordFreqHeap, wordFreq{Word: word, Count: count})
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(&minwordFreqHeap).(wordFreq).Word
	}
	return res
}

func Test_topKFrequent() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))                        //[i love]
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4)) //[the is sunny day]
}
