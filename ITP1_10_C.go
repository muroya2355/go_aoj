package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	for {
		fmt.Scan(&n)
		if n==0 {
			break
		}
		s := make([]float64, n)
		for i:=0; i<n; i++ {
			fmt.Scan(&s[i])
		}

		mean := 0.0
		for i:=0; i<n; i++ {
			mean += s[i]/float64(n)
		}
		sdv := 0.0
		for i:=0; i<n; i++ {
			sdv += (s[i]-mean)*(s[i]-mean) / float64(n)
		}

		fmt.Println(math.Sqrt(sdv))
	}

}