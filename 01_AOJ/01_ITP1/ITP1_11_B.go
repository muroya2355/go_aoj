package main

import "fmt"

type Dice struct {
	side [6]int
	topnum int
	rightnum, leftnum, forwardnum, backnum, bottomnum int
}
func (dice *Dice) init() {
	for i:=0; i<6; i++ {
		fmt.Scan(&(dice.side[i]))
	}
	dice.topnum, dice.rightnum, dice.leftnum, dice.forwardnum, dice.backnum, dice.bottomnum = 0, 2, 3, 4, 1, 5
}
func (dice *Dice) dtonum(d int) int {
	for i:=0; i<6; i++ {
		if dice.side[i]==d {
			return i
		}
	}
	return 0
}
func (dice *Dice) inittopback(top, back int) {
	//for i:=0; i<6; i++ {
	//	fmt.Scan(&(dice.side[i]))
	//}
	dice.topnum, dice.rightnum, dice.leftnum, dice.forwardnum, dice.backnum, dice.bottomnum = top, 2, 3, 5-back, back, 5-top
	topbackpairs := [6][4]int{{3,1,2,4},{0,3,5,2},{0,1,5,4},{1,0,4,5},{2,5,3,0},{2,1,3,4}}
	for i:=0; i<6; i++ {
		for j:=0; j<4; j++ {
			if top==topbackpairs[i][j] && back==topbackpairs[i][(j+1)%4] {
				dice.rightnum = i
				dice.leftnum = 5-i
				break
			}
		}
	}

}
func (dice *Dice) rotate(dir string) {
	temp := dice.topnum

	switch dir {
	case "N" :	
		dice.topnum = dice.backnum
		dice.backnum = dice.bottomnum
		dice.bottomnum = dice.forwardnum
		dice.forwardnum = temp
	case "E" :
		dice.topnum = dice.leftnum
		dice.leftnum = dice.bottomnum
		dice.bottomnum = dice.rightnum
		dice.rightnum = temp
	case "W" :
		dice.topnum = dice.rightnum
		dice.rightnum = dice.bottomnum
		dice.bottomnum = dice.leftnum
		dice.leftnum = temp
	case "S" :
		dice.topnum = dice.forwardnum
		dice.forwardnum = dice.bottomnum
		dice.bottomnum = dice.backnum
		dice.backnum = temp
	}
}

func main() {

	var dice Dice
	dice.init()
	var q int
	fmt.Scan(&q)
	for i:=0; i<q; i++ {
		var top, back int
		fmt.Scan(&top, &back)
		dice.inittopback(dice.dtonum(top), dice.dtonum(back))
		fmt.Println(dice.side[dice.rightnum])
	}
	  
}