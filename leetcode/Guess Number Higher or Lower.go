package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	return bs_guessNumber(1, n)
}

func bs_guessNumber(left, right int) int {
	mid := left + (right-left)/2
	res := guess(mid)
	if res == 0 {
		return mid
	} else if res == 1 {
		return bs_guessNumber(mid+1, right)
	} else {
		return bs_guessNumber(left, mid)
	}
}

// only for test
func guess(num int) int {
	return 0
}
