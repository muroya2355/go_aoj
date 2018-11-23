package main

import "fmt"

func main() {

	var n, x int

	for {
		fmt.Scan(&n, &x)
		if n==0 && x==0 {
			break
		}
		tot := 0
		for a:=1; a<=n-2; a++ {
			for b:=a+1; b<=n-1; b++ {
				for c:=b+1; c<=n; c++ {
					if a+b+c == x {
						tot++
					}
				}
			}
		}
		fmt.Println(tot)
	}
}