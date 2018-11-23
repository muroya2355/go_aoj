package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scan(&n)
	x := make([]float64, n)
	y := make([]float64, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&x[i])
	}
	for i:=0; i<n; i++ {
		fmt.Scan(&y[i])
	}

	dis1, dis2, dis3, disInf := 0.0, 0.0, 0.0, 0.0

	for i:=0; i<n; i++ {
		abs := math.Abs(x[i]-y[i])
		dis1 += abs
		dis2 += abs*abs
		dis3 += abs*abs*abs
		disInf = math.Max(abs,disInf)
	}

	fmt.Println(dis1)
	fmt.Println(math.Pow(dis2,1.0/2.0))
	fmt.Println(math.Pow(dis3,1.0/3.0))
	fmt.Println(disInf)

}