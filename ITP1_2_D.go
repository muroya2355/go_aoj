package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	var w, h, x, y, r int
	fmt.Scan(&w, &h, &x, &y, &r)

	var top, bottom, left, right Vertex
	top 	= Vertex{x,y+r}
	bottom 	= Vertex{x,y-r}
	left 	= Vertex{x-r,y}
	right	= Vertex{x+r,y}

	if left.X >= 0 && top.Y <= h && right.X <= w && bottom.Y >= 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}