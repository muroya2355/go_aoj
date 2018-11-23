package main

import "fmt"

func print(rstr []rune, a,b int) {
	
	for i:=a; i<=b; i++ {
		fmt.Print(string(rstr[i]))
	}
	fmt.Println()
}
func swap(x,y rune)(rune, rune) {
	return y, x
}
func reverse(rstr []rune, a,b int) []rune {
	
	for i:=0; i<(b-a+1)/2; i++ {
		rstr[a+i], rstr[b-i] = swap(rstr[a+i], rstr[b-i])
	}
	return rstr
}
func replace(rstr []rune, p []rune, a,b int) []rune {

	for i:=a; i<=b; i++ {
		rstr[i] = p[i-a]
	}
	return rstr
}

func main() {

	var str string
	fmt.Scan(&str)
	//rstr := []rune(str)
	var n int
	fmt.Scan(&n)

	for i:=0; i<n; i++ {
		var op, p string
		var a, b int
		fmt.Scan(&op, &a, &b)

		if op=="replace" {
			fmt.Scan(&p)
			str = string(replace([]rune(str),[]rune(p),a,b))
		}
		if op=="reverse" {
			str = string(reverse([]rune(str),a,b))
		}
		if op=="print" {
			print([]rune(str),a,b)
		}
	} 
}