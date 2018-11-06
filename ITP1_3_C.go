package main

import "fmt"

func sort(x, y int)(int, int) {
	if x > y {
		return y, x
	}
	return x, y
}
func main() {

	var a, b int
	fmt.Scan(&a, &b)
	for a != 0 || b != 0 {
		a, b = sort(a, b)
		fmt.Println(a, b)
		fmt.Scan(&a, &b)
	}
}