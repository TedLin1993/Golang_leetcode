package leetcode

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	letterList := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	wordMap := map[string]int{}
	for _, word := range words {
		morseWord := ""
		for i := 0; i < len(word); i++ {
			morseWord += letterList[word[i]-'a']
		}
		wordMap[morseWord]++
	}
	return len(wordMap)
}

func Test_uniqueMorseRepresentations() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})) //2
	fmt.Println(uniqueMorseRepresentations([]string{"a"}))                        //1
}
