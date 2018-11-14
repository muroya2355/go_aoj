package main

import "fmt"

func main() {

	var r, c, a_ij, tot int
	tot = 0

	fmt.Scan(&r,&c)
	a := make([][]int, r)
	for i:=0; i<r; i++ {
		a[i] = make([]int, c)
	}
	rsum := make([]int, r)
	csum := make([]int, c)

	for i:=0; i<r; i++ {
		for j:=0; j<c; j++ {
			fmt.Scan(&a_ij)
			a[i][j] = a_ij
			rsum[i] += a_ij
			csum[j] += a_ij
			tot += a_ij
		}
	}
	for i:=0; i<r; i++ {
		for j:=0; j<c; j++ {
			fmt.Print(a[i][j]," ")
		}
		fmt.Print(rsum[i],"\n")
	}
	for j:=0; j<c; j++ {
		fmt.Print(csum[j]," ")
	}
	fmt.Print(tot,"\n")
}