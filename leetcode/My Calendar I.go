package leetcode

import (
	"fmt"
)

type MyCalendar struct {
	calendars [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.calendars) == 0 {
		this.calendars = append(this.calendars, []int{start, end})
		return true
	}

	if start < this.calendars[0][0] {
		if end <= this.calendars[0][0] {
			this.calendars = append([][]int{{start, end}}, this.calendars...)
			return true
		}
		return false
	}
	if start > this.calendars[len(this.calendars)-1][0] {
		if start >= this.calendars[len(this.calendars)-1][1] {
			this.calendars = append(this.calendars, []int{start, end})
			return true
		}
		return false
	}

	left, right := 1, len(this.calendars)-1
	for left < right {
		mid := (right + left + 1) / 2
		if this.calendars[mid-1][0] > start {
			right = mid - 1
		} else {
			left = mid
		}
	}
	index := left
	if this.calendars[index-1][1] > start || this.calendars[index][0] < end {
		return false
	}
	this.calendars = append(this.calendars[:index+1], this.calendars[index:]...)
	this.calendars[index] = []int{start, end}
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
func TestMyCalendar() {
	obj := Constructor()
	fmt.Println(obj.Book(20, 29)) //true
	fmt.Println(obj.Book(13, 22)) //false
	fmt.Println(obj.Book(44, 50)) //true
	fmt.Println(obj.Book(1, 7))   //true
	fmt.Println(obj.Book(2, 10))  //false
	fmt.Println(obj.Book(14, 20)) //true
	fmt.Println(obj.Book(19, 25)) //false
	fmt.Println(obj.Book(36, 42)) //true
	// fmt.Println(obj.Book(8, 13))  //true
	// fmt.Println(obj.Book(18, 27)) //false
}
