package leetcode

import "container/heap"

type charFreq struct {
	Char  byte
	Count int
}

func reorganizeString(s string) string {
	n := len(s)
	charMap := make(map[byte]int, n)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}
	maxCfHeap := MaxCFHeap{}
	for char, count := range charMap {
		heap.Push(&maxCfHeap, charFreq{char, count})
	}

	res := make([]byte, 0, n)
	for len(maxCfHeap) >= 2 {
		cf1 := heap.Pop(&maxCfHeap).(charFreq)
		cf2 := heap.Pop(&maxCfHeap).(charFreq)
		res = append(res, cf1.Char, cf2.Char)
		cf1.Count--
		if cf1.Count > 0 {
			heap.Push(&maxCfHeap, cf1)
		}
		cf2.Count--
		if cf2.Count > 0 {
			heap.Push(&maxCfHeap, cf2)
		}
	}
	if len(maxCfHeap) == 1 {
		lastCf := maxCfHeap[0]
		if lastCf.Count > 1 {
			return ""
		}
		res = append(res, lastCf.Char)
	}

	return string(res)
}

type MaxCFHeap []charFreq

func (h MaxCFHeap) Len() int           { return len(h) }
func (h MaxCFHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h MaxCFHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxCFHeap) Push(x any) {
	*h = append(*h, x.(charFreq))
}
func (h *MaxCFHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
