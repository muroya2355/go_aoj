package main

import "fmt"

func main() {

	var n, m, a_ij, b_j int
	fmt.Scan(&n, &m)
	
	a := make([][]int, n)
	for i:=0; i<n; i++ {
		a[i] = make([]int, m)
		for j:=0; j<m; j++ {
			fmt.Scan(&a_ij)
			a[i][j] = a_ij
		}
	}
	b := make([]int, m)
	for j:=0; j<m; j++ {
		fmt.Scan(&b_j)
		b[j] = b_j
	}

	c := make([]int, n)
	for i:=0; i<n; i++ {
		c[i] = 0
		for j:=0; j<m; j++ {
			c[i] += a[i][j]*b[j]
		}
		fmt.Println(c[i])
	}

}