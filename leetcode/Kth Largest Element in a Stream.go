package leetcode

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	k          int
	sortedNums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{k: k, sortedNums: nums}
}

func (this *KthLargest) Add(val int) int {
	i := sort.SearchInts(this.sortedNums, val)
	if i == len(this.sortedNums) {
		this.sortedNums = append(this.sortedNums, val)
	} else {
		this.sortedNums = append(this.sortedNums[:i+1], this.sortedNums[i:]...)
		this.sortedNums[i] = val
	}
	this.sortedNums = this.sortedNums[len(this.sortedNums)-this.k:]
	return this.sortedNums[0]
}

func TestKthLargest() {
	var kthLargest KthLargest
	//test 1
	kthLargest = Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kthLargest.Add(3))  //4
	fmt.Println(kthLargest.Add(5))  //5
	fmt.Println(kthLargest.Add(10)) //5
	fmt.Println(kthLargest.Add(9))  //8
	fmt.Println(kthLargest.Add(4))  //8

	//test 2
	kthLargest = Constructor(1, []int{})
	fmt.Println(kthLargest.Add(-3)) //-3
	fmt.Println(kthLargest.Add(-2)) //-2
	fmt.Println(kthLargest.Add(-4)) //-2
	fmt.Println(kthLargest.Add(0))  //0
	fmt.Println(kthLargest.Add(4))  //4
}
