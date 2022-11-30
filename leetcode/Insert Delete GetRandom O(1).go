package leetcode

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	valMap  map[int]int
	valList []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{valMap: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.valMap[val]; exist {
		return false
	}
	this.valList = append(this.valList, val)
	this.valMap[val] = len(this.valList) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, exist := this.valMap[val]; !exist {
		return false
	}
	lastVal := this.valList[len(this.valList)-1]
	idx := this.valMap[val]
	this.valList[idx] = lastVal
	this.valMap[lastVal] = idx

	this.valList = this.valList[:len(this.valList)-1]
	delete(this.valMap, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.valList))
	return this.valList[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func Test_RandomizedSet() {
	obj := Constructor()
	fmt.Println(obj.Remove(1))   //false
	fmt.Println(obj.Insert(1))   //true
	fmt.Println(obj.GetRandom()) //1
	fmt.Println(obj.Remove(1))   //true
	fmt.Println(obj.Insert(1))   //true
	fmt.Println(obj.Insert(1))   //false
	fmt.Println(obj.Insert(2))   //true
	fmt.Println(obj.GetRandom()) //1 or 2
	fmt.Println(obj.GetRandom()) //1 or 2
	fmt.Println(obj.Remove(1))   //true
	fmt.Println(obj.Remove(1))   //false
}
