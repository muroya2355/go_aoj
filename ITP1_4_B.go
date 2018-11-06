package main

import (
	"fmt"
	"math"
)

func main() {

	var r float64
	fmt.Scan(&r)
	len := 2*r*math.Pi
	s := math.Pi*r*r
	fmt.Printf("%.6f %.6f", s, len)
}