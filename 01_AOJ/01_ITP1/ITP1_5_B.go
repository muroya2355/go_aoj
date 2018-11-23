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
			fmt.Print("#")
		
			for col:=1; col<w-1; col++ {
				if row==0 || row==h-1 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("#\n")
		}
		fmt.Print("\n")
	}
}