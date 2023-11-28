package leetcode

func beautifulSubstrings(s string, k int) int {
	n := len(s)
	vowelMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	res := 0
	for i := 0; i < n; i++ {
		vowels, consonants := 0, 0
		for j := i; j < n; j++ {
			if vowelMap[s[j]] {
				vowels++
			} else {
				consonants++
			}
			if vowels == consonants && vowels*consonants%k == 0 {
				res++
			}
		}
	}
	return res
}
