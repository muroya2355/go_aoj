package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Hands [3]string = [3]string{"goo", "choki", "parm"}

type CP struct {
	Hand string
	Win int
	Lose int
	Even int
}

func (cp *CP) dohands() {
	rand.Seed(time.Now().UnixNano())
	cp.Hand = Hands[rand.Intn(3)]
}

func (cp1 *CP) rspgo(cp2 CP) {
	if cp1.Hand == cp2.Hand {
		cp1.Even++
	} else if cp1.Hand=="goo" && cp2.Hand=="choki" {
		cp1.Win++
	} else if cp1.Hand=="goo" && cp2.Hand=="parm" {
		cp1.Lose++
	} else if cp1.Hand=="choki" && cp2.Hand=="goo" {
		cp1.Lose++
	} else if cp1.Hand=="choki" && cp2.Hand=="parm" {
		cp1.Win++
	} else if cp1.Hand=="parm" && cp2.Hand=="goo" {
		cp1.Win++
	} else if cp1.Hand=="parm" && cp2.Hand=="choki" {
		cp1.Lose++
	}
}

func main() {

	cp1 := CP{Hands[0],0,0,0}
	cp2 := CP{Hands[0],0,0,0}
	
	for i:=0; i<1000000; i++ {
		cp1.dohands()
		cp2.dohands()
		cp1.rspgo(cp2)
	}

	fmt.Println("勝ち： ",cp1.Win)
	fmt.Println("相子： ",cp1.Even)
	fmt.Println("負け： ",cp1.Lose)
}