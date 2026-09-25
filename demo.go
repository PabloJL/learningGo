package main

import (
	"fmt"
	"math"
)

func Score(x, y float64) int {
	var score int
	a:= x * x
	b:= y * y
	distance := math.Sqrt(a + b)
	
	if distance <= 1 {
		score = 10
	}else if distance <= 5{
		score = 5
	}else if distance <= 10{
		score = 1
	}else{
		score = 0
	}
	return score
}

func main() {
	fmt.Println(Score(1,0))
}

