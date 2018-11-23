package main

import "fmt"

func main() {

	var i, x int
	i = 1
	fmt.Scan(&x)
	for x != 0 {
		fmt.Printf("Case %d: %d\n", i, x)
		fmt.Scan(&x)
		i++
	}
}