package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x = x + Calc(x)
		fmt.Println(x)
	}
	fmt.Println("Code.education Rocks!")
}

func Calc(number float64) float64 {
	return math.Sqrt(number)
}
