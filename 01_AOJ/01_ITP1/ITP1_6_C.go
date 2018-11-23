package main

import "fmt"

func main() {

	ohouse := make([][][]int, 4)
	for building:=0; building<4; building++ {
		ohouse[building] = make([][]int, 3)
		for floor:=0; floor<3; floor++ {
			ohouse[building][floor] = make([]int, 10)
			for room:=0; room<10; room++ {
				ohouse[building][floor][room] = 0
			}
		}
	}

	var n, b, f, r, v int
	fmt.Scan(&n)
	for m:=0; m<n; m++ {
		fmt.Scan(&b, &f, &r, &v)
		ohouse[b-1][f-1][r-1] += v
	}

	for building:=0; building<4; building++ {
		for floor:=0; floor<3; floor++ {
			for room:=0; room<10; room++ {
				fmt.Print(" ",ohouse[building][floor][room])
			}
			fmt.Println()
		}
		if building != 3 {
			fmt.Println("####################")
		}
	}
}