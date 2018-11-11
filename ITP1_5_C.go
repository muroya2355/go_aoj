package main

import "fmt"

func main() {

	var h, w int
	
	for {

		fmt.Scan(&h, &w)
		if h==0 && w==0 {
			break
		}

		for row:=0; row<h; row++ {
			for col:=0; col<w; col++ {
				if (row+col)%2 == 0{
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}