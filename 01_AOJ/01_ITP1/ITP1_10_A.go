package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func Distance(a,b Point) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)) 
}

func main() {

	var a, b Point
	fmt.Scan(&(a.X),&(a.Y),&(b.X),&(b.Y))

	fmt.Println(Distance(a,b))
}