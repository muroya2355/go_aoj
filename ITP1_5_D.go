package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	for i:=1; i<=n; i++ {
		
		x := i
		if (x%3)==0 {
			fmt.Printf(" %d",i) 
			continue
		}
		for (x%10)==3 || (x/10)!=0 {
			if (x%10)==3 {
				fmt.Printf(" %d",i) 
				break
			} else if (x/10)!=0 {
				x = x/10
			}
		}
	}
	fmt.Println()
}