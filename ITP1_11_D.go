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
func (dice1 *Dice) checksameside(dice2 Dice) bool {

	if dice1.side[dice1.topnum] == dice2.side[dice2.topnum] && dice1.side[dice1.leftnum] == dice2.side[dice2.leftnum] && dice1.side[dice1.rightnum] == dice2.side[dice2.rightnum] && dice1.side[dice1.forwardnum] == dice2.side[dice2.forwardnum] && dice1.side[dice1.backnum] == dice2.side[dice2.backnum] && dice1.side[dice1.bottomnum] == dice2.side[dice2.bottomnum] {
			
		return true
	} else {
		return false
	}
}
func (dice1 Dice) havesameside(dice2 Dice) bool {
	
	for i:=0; i<4; i++ {
		if dice1.checksameside(dice2) {
			return true
		} else {
			for j:=0; j<4; j++ {
				dice1.rotate("E")
				if dice1.checksameside(dice2) {
					return true
				}
			}
		}
		dice1.rotate("N")
	}
	dice1.rotate("E")
	for i:=0; i<4; i++ {
		if dice1.checksameside(dice2) {
			return true
		} else {
			for j:=0; j<4; j++ {
				dice1.rotate("E")
				if dice1.checksameside(dice2) {
					return true
				}
			}
		}
		dice1.rotate("N")
	}
	return false
}

func main() {

	var n int
	fmt.Scan(&n)
	dices := make([]Dice, n)
	for i:=0; i<n; i++ {
		dices[i].init()
	}

	differentdices := true
	for i:=0; i<n-1; i++ {
		for j:=i+1; j<n; j++ {
			differentdices = differentdices && !(dices[i].havesameside(dices[j]))
		}
	}
	
	if differentdices {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}