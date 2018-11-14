package main

import "fmt"

func main() {

	var n, m, l, a_ij, b_jk int
	fmt.Scan(&n, &m, &l)
	
	a := make([][]int, n)
	for i:=0; i<n; i++ {
		a[i] = make([]int, m)
		for j:=0; j<m; j++ {
			fmt.Scan(&a_ij)
			a[i][j] = a_ij
		}
	}
	b := make([][]int, m)
	for j:=0; j<m; j++ {
		b[j] = make([]int, l)
		for k:=0; k<l; k++ {
			fmt.Scan(&b_jk)
			b[j][k] = b_jk
		}
	}

	c := make([][]int, n)
	for i:=0; i<n; i++ {
		c[i] = make([]int, l)
		for k:=0; k<l; k++ {
			for j:=0; j<m; j++ {
				c[i][k] += a[i][j]*b[j][k]
			}
			if k==l-1 {
				fmt.Print(c[i][k])
			} else {
				fmt.Print(c[i][k]," ")
			}
		}
		fmt.Println()
	}

}