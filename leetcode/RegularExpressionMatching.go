package leetcode

import "fmt"

// func isMatch(s string, p string) bool {
// 	s_index, p_index := 0, 0
// 	for s_index < len(s) && p_index < len(p) {
// 		if isTheSame(s[s_index], p[p_index]) {
// 			s_index++
// 			p_index++
// 		} else {
// 			if p[p_index] == '*' {
// 				if isTheSame(s[s_index], p[p_index-1]) {
// 					//check previous char
// 					s_index++
// 				} else if p_index+1 < len(p) && isTheSame(s[s_index], p[p_index+1]) {
// 					//check next char
// 					s_index++
// 					p_index = p_index + 2
// 				} else {
// 					return false
// 				}
// 			} else if p_index+1 < len(p) && p[p_index+1] == '*' {
// 				//check next char
// 				s_index++
// 				p_index = p_index + 2
// 			} else {
// 				return false
// 			}
// 		}
// 	}
// 	if s_index == len(s) {
// 		if p_index == len(p) {
// 			return true
// 		}
// 		if p_index == len(p)-1 && p[p_index] == '*' {
// 			return true
// 		}
// 	}
// 	return false
// }

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 {
		return len(s) == 1 && isTheSame(s[0], p[0])
	}
	if len(p) == 2 && len(s) == 0 {
		return p[1] == '*'
	}
	if len(s) == 0 {
		return false
	}
	if p[1] != '*' {
		return isTheSame(s[0], p[0]) && isMatch(s[1:], p[1:])
	} else {
		if isTheSame(s[0], p[0]) {
			if len(s) > 1 && (s[1] == s[0] || p[0] == '.') {
				return isMatch(s[1:], p) || isMatch(s, p[2:]) || isMatch(s[1:], p[2:])
			} else {
				return isMatch(s, p[2:]) || isMatch(s[1:], p[2:])
			}
		} else {
			return isMatch(s, p[2:])
		}
	}
}

func isTheSame(a, b byte) bool {
	if a == b || b == '.' {
		return true
	}
	return false
}

func TestRegularExpressionMatching() {
	// fmt.Println(isMatch("aa", "a"))          //false
	// fmt.Println(isMatch("aaa", "ab*a"))      //false
	// fmt.Println(isMatch("abcd", "d*"))       //false
	// fmt.Println(isMatch("sippi", "s*p*."))   //false
	// fmt.Println(isMatch("ippi", "s*p*."))    //false
	// fmt.Println(isMatch("aa", "a*"))         //true
	// fmt.Println(isMatch("ab", ".*"))         //true
	// fmt.Println(isMatch("abcasdas", ".*"))   //true
	// fmt.Println(isMatch("aab", "c*a*b"))     //true
	// fmt.Println(isMatch("aaa", "a*a"))       //true
	// fmt.Println(isMatch("aaaaaaaa", "a*aa")) //true
	// fmt.Println(isMatch("aaab", "a*ab"))     //true
	// fmt.Println(isMatch("aa", "a*b*"))       //true
	// fmt.Println(isMatch("ba", ".*a*a"))      //true
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*b*b*")) //true

}
