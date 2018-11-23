package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, C float64
	var S, L, h float64

	fmt.Scan(&a, &b, &C)

	theta := C/180.0 * math.Pi
	S = 0.5 * a * b * math.Sin(theta)
	L = math.Sqrt(a*a + b*b - 2*a*b*math.Cos(theta)) + a + b
	h = b*math.Sin(theta)

	fmt.Println(S)
	fmt.Println(L)
	fmt.Println(h)
}