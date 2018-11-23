package main

import "fmt"

func main() {

	var n, min, max, sum, temp int
	fmt.Scan(&n)
	fmt.Scan(&temp)
	min = temp
	max = temp
	sum = temp

	for i:=1; i<n; i++ {
		fmt.Scan(&temp)
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}
		sum += temp
	}
	fmt.Println(min, max, sum)
}
