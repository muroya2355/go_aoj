package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {

	A := make([]int, 200010)
	var n int
	fmt.Scan(&n)
	
	last_idx := -1
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i:=0; i<n; i++ {
		
		scanner.Scan()
		query, _ := strconv.Atoi(scanner.Text())

		switch query {
		case 0 :
			last_idx++
			scanner.Scan()
			x, _ := strconv.Atoi(scanner.Text())
			A[last_idx] = x
		case 1 :
			scanner.Scan()
			p, _ := strconv.Atoi(scanner.Text())
			fmt.Println(A[p])
		case 2 :
			A[last_idx] = 0
			last_idx--
		}	
	}
}
