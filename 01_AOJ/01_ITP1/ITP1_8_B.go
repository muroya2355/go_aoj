package main

import(
	"fmt"
)

func main() {

	var numstr string
	
	for {
		fmt.Scan(&numstr)
		if numstr=="0" {
			break
		}

		rnumstr := []rune(numstr)

		tot := 0
		for _, rnumchar := range rnumstr {
			//fmt.Println(string(rnumchar))
			tot += int(rnumchar) - int(rune('0'))
		}
		fmt.Println(tot)
	}
}