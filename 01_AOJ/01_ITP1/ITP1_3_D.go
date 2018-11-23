package main

import "fmt"

func min(x, y int) int{
	if x < y {
		return x
	}
	return y
}

func main() {

	var a, b, c, n int
	fmt.Scan(&a, &b, &c)
	n = 0

	for i := a; i <= b; i++ {
		if c % i == 0 {
			n++
		}
	}
	fmt.Println(n)
}