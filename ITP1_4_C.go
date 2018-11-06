package main

import "fmt"

func main() {

	var x, y int
	var op string
	fmt.Scan(&x, &op, &y)

	for op!="?" {
		switch op {
		case "+":
			fmt.Println(x+y)
		case "-":
			fmt.Println(x-y)
		case "*":
			fmt.Println(x*y)
		case "/":
			fmt.Println(x/y)
		}
		fmt.Scan(&x,&op,&y)
	}
}