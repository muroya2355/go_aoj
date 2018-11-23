package main

import "fmt"

func main() {

	var a, b int64
	fmt.Scan(&a,&b)
	fmt.Printf("%d %d %.9f\n", a/b, a%b, float64(a)/float64(b))
}