package main

import "fmt"

func main() {

	var x, y int

	for  {
		fmt.Scan(&y,&x)
		if x==0 && y==0 {
			break
		}

		for row:=0; row<y; row++ {
			for col:=0; col<x; col++ {
				fmt.Print("#")
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}