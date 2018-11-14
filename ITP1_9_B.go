package main

import "fmt"

func main() {

	var str string

	for {
		fmt.Scan(&str)
		if str=="-" {
			break
		}

		var n, m int
		rstr := []rune(str)
		n = len(rstr)
		rstr_temp := make([]rune, n)
		fmt.Scan(&m)

		for i:=0; i<m; i++ {
			
			var h int
			fmt.Scan(&h)
			_ = copy(rstr_temp, rstr)

			for idx:=0; idx<n; idx++ {
				rstr[idx] = rstr_temp[(idx+h)%n]
			}
		}
		fmt.Println(string(rstr))
	}
}