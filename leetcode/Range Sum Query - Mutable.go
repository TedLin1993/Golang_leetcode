package leetcode

type NumArray struct {
	bit  []int
	nums []int
}

func Constructor_NumArray(nums []int) NumArray {
	bit := make([]int, len(nums)+1)
	for i, v := range nums {
		updateBit(&bit, i, v)
	}
	return NumArray{bit: bit, nums: nums}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.nums[index]
	this.nums[index] = val
	updateBit(&this.bit, index, diff)
}

func updateBit(bit *[]int, index int, val int) {
	index++
	for index < len(*bit) {
		(*bit)[index] += val
		index += index & -index
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum(right) - this.sum(left-1)
}

func (this *NumArray) sum(index int) int {
	index++
	res := 0
	for index > 0 {
		res += this.bit[index]
		index -= index & -index
	}
	return res
}
