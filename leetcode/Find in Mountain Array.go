package leetcode

/**
 * This is the MountainArray's API interface.
 * You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

// only for testing purposes
type MountainArray struct{}

func (this *MountainArray) get(index int) int { return 0 }
func (this *MountainArray) length() int       { return 0 }

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	peak := left
	left, right = 0, peak
	for left <= right {
		mid := left + (right-left)/2
		midVal := mountainArr.get(mid)
		if midVal > target {
			right = mid - 1
		} else if midVal < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	left, right = peak, n-1
	for left <= right {
		mid := left + (right-left)/2
		midVal := mountainArr.get(mid)
		if midVal > target {
			left = mid + 1
		} else if midVal < target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
