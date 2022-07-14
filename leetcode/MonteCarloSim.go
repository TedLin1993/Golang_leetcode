package leetcode

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func monteCarloSim(count int) {
	rand.Seed(time.Now().UnixNano())
	originCount := 0
	rest := 0
	for i := 0; i < count; i++ {
		x, y, z := rand.Float64(), rand.Float64(), rand.Float64()
		distBigBall := math.Sqrt(x*x + y*y + z*z)
		distXBall := math.Sqrt(math.Pow((x-0.5), 2) + y*y + z*z)
		distYBall := math.Sqrt(x*x + math.Pow((y-0.5), 2) + z*z)
		distZBall := math.Sqrt(x*x + y*y + math.Pow((z-0.5), 2))
		if distBigBall <= 1 {
			originCount++
			if distXBall > 0.5 && distYBall > 0.5 && distZBall > 0.5 {
				rest++
			}
		}
	}

	fmt.Printf("origin: %v, rest: %v, rest/origin: %v \n", originCount, rest, float64(rest)/float64(originCount))
}

func TestmonteCarloSim() {
	monteCarloSim(1000000)
	monteCarloSim(1000000)
	monteCarloSim(1000000)
}
