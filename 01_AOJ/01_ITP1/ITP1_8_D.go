package main

import "fmt"

func main() {

	var s, p string
	fmt.Scan(&s,&p)
	s += s
	rs := []rune(s)
	rp := []rune(p)

	for i:=0; i<len(rs)/2+1; i++ {
		flag := 1
		for j:=0; j<len(rp); j++ {
			if rp[j]==rs[i+j] {
				flag *= 1
			} else {
				flag *= 0
			}
		}
		if flag==1 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
	return
}