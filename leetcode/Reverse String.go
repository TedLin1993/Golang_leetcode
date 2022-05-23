package leetcode

import "fmt"

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func TestreverseString() {
	var s []byte
	s = []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s) //[111 108 108 101 104]
	s = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)
	fmt.Println(s) //[104 97 110 110 97 72]
}
