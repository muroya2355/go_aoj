package main

import "fmt"

func main() {

	deck := make(map[string][]int)
	shcd := [4]string{"S", "H", "C", "D"}

	for i:=0; i<4; i++ {
		deck[shcd[i]] = make([]int,13)
	}

	var n, num int
	var mark string
	fmt.Scan(&n)

	for m:=0; m<n; m++ {
		fmt.Scan(&mark, &num)
		deck[mark][num-1] = 1
	}
	for i:=0; i<4; i++ {
		for j:=0; j<13; j++{
			if deck[shcd[i]][j]==0 {
				fmt.Println(shcd[i],j+1)
			}
		}
	}

}