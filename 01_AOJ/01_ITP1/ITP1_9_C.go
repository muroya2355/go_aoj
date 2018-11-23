package main

import "fmt"

func main() {

	pTaro := 0
	pHanako := 0
	var n int
	fmt.Scan(&n)

	for i:=0; i<n; i++ {
		var cTaro, cHanako string
		fmt.Scan(&cTaro, &cHanako)
		if cTaro > cHanako {
			pTaro += 3
		} else if cTaro < cHanako {
			pHanako += 3
		} else {
			pTaro += 1
			pHanako += 1
		}
	}
	fmt.Println(pTaro, pHanako)
}