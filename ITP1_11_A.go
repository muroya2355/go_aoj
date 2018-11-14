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
	var operator string
	fmt.Scan(&operator)
	roperator := []rune(operator)
	for i:=0; i<len(roperator); i++ {
		dice.rotate(string(roperator[i]))
	}
	fmt.Println(dice.side[dice.topnum])
	  
}