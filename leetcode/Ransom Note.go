package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteLetters := make([]int, 26)
	for i := 0; i < len(ransomNote); i++ {
		ransomNoteLetters[ransomNote[i]-'a']++
	}

	magazineLetters := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		magazineLetters[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNoteLetters); i++ {
		if ransomNoteLetters[i] > magazineLetters[i] {
			return false
		}
	}
	return true
}
