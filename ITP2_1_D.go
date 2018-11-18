package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {

	var n, q int
	fmt.Scan(&n, &q)
	A := make([][]string, n)
	for t:=0; t<n; t++ {
		A[t] = append(A[t],"NULL")
	}
	
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	//var writer = bufio.NewWriter(os.Stdout)

	for i:=0; i<q; i++ {
		
		scanner.Scan()
		query, _ := strconv.Atoi(scanner.Text())

		switch query {
		case 0 :
			scanner.Scan()
			t, _ := strconv.Atoi(scanner.Text())
			scanner.Scan()
			x := scanner.Text()
			if A[t][0] == "NULL" {
				A[t][0] = x
			} else {
				A[t] = append(A[t], x)
			}
		case 1 :
			scanner.Scan()
			t, _ := strconv.Atoi(scanner.Text())
			if A[t][0] != "NULL" {
				for idx:=0; idx<len(A[t]); idx++ {
					if idx==len(A[t])-1 {
						//fmt.Fprintln(writer, A[t][idx])
						fmt.Println(A[t][idx])
					} else {
						//fmt.Fprintf(writer, A[t][idx], " ")
						fmt.Print(A[t][idx], " ")
					}
				}
				//_ = writer.Flush()
			} else {
				fmt.Println()
			}
		case 2 :
			scanner.Scan()
			t, _ := strconv.Atoi(scanner.Text())
			A[t] = A[t][0:1]
			A[t][0] = "NULL"
		}	
	}
}
