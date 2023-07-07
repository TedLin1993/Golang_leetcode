package leetcode

func findPrimePairs(n int) [][]int {
	var primes []int
	var isPrime []bool
	getAllPrime := func(max int) {
		isPrime = make([]bool, max+1)
		for i := range isPrime {
			isPrime[i] = true
		}
		for num := 2; num < max; num++ {
			if isPrime[num] {
				primes = append(primes, num)
				num2 := num * num
				for num2 < max {
					isPrime[num2] = false
					num2 += num
				}
			}
		}
	}
	getAllPrime(n)
	res := [][]int{}
	for i := 0; primes[i] <= n && primes[i] <= n/2; i++ {
		if isPrime[n-primes[i]] {
			res = append(res, []int{primes[i], n - primes[i]})
		}
	}
	return res
}
