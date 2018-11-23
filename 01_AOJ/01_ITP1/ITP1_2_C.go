package main

import "fmt"

func swap(x, y int)(int, int){
	if x > y {
		return y, x
	}
	return x, y
}

func main() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)
	
	a, b = swap(a, b)
	a, c = swap(a, c)
	b, c = swap(b, c)

	fmt.Println(a, b, c)
}