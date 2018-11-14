package main

import "fmt"

func main() {

	var n, a_i int
	var buf [100]int

	fmt.Scan(&n)
	var a []int = buf[0:n]
	for i:=0; i<n; i++ {
		fmt.Scan(&a_i)
		a[i] = a_i
	}
	for i:=n-1; i>0; i-- {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println(a[0])
	a = nil
}