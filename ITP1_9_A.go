package main

import (
	"fmt"
	"strings"
)

func main() {

	var t, w string
	var n int = 0
	fmt.Scan(&w)

	for {
		fmt.Scan(&t)
		if t=="END_OF_TEXT" {
			break
		}
		t = strings.ToLower(t)
		if t==w {
			n++
		}
	}
	fmt.Println(n)
}