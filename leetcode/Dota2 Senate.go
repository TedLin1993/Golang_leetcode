package leetcode

import "fmt"

func predictPartyVictory(senate string) string {
	rCount, dCount := 0, 0
	banR, banD := 0, 0
	banned := make([]bool, len(senate))
	for {
		for i := 0; i < len(senate); i++ {
			if banned[i] {
				continue
			}
			if senate[i] == 'R' {
				if banR > 0 {
					banned[i] = true
					banR--
				} else {
					rCount++
					banD++
				}
			} else {
				if banD > 0 {
					banned[i] = true
					banD--
				} else {
					dCount++
					banR++
				}
			}
		}
		if rCount == 0 || dCount == 0 {
			break
		}
		rCount, dCount = 0, 0
	}

	if dCount > 0 {
		return "Dire"
	}
	return "Radiant"
}

func predictPartyVictory_2(senate string) string {
	n := len(senate)
	rqueue, dqueue := make([]int, 0, n), make([]int, 0, n)
	for idx, char := range senate {
		if char == 'R' {
			rqueue = append(rqueue, idx)
		} else {
			dqueue = append(dqueue, idx)
		}
	}

	for len(rqueue) > 0 && len(dqueue) > 0 {
		r, d := rqueue[0], dqueue[0]
		rqueue = rqueue[1:]
		dqueue = dqueue[1:]
		if r < d {
			rqueue = append(rqueue, r+n)
		} else {
			dqueue = append(dqueue, d+n)
		}
	}
	if len(rqueue) > 0 {
		return "Radiant"
	}
	return "Dire"
}

func Test_predictPartyVictory() {
	fmt.Println(predictPartyVictory_2("RD"))    //R
	fmt.Println(predictPartyVictory_2("RDD"))   //D
	fmt.Println(predictPartyVictory_2("DR"))    //D
	fmt.Println(predictPartyVictory_2("DDRRR")) //D
	fmt.Println(predictPartyVictory_2("DRRD"))  //D
}
