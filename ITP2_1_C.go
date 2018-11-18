package main

import (
	"container/list"
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {

	L := list.New()
	L.PushBack(-100)
	cur := L.Front()
	var n int
	fmt.Scan(&n)
	
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i:=0; i<n; i++ {
		
		scanner.Scan()
		query, _ := strconv.Atoi(scanner.Text())

		switch query {
		case 0 :
			scanner.Scan()
			x, _ := strconv.Atoi(scanner.Text())
			if i==0 {
				L.PushFront(x)
				cur = L.Front()
			} else {
				L.InsertBefore(x, cur)
				cur = cur.Prev()
			}
		case 1 :
			scanner.Scan()
			d, _ := strconv.Atoi(scanner.Text())
			if d > 0 {
				for j:=0; (j<d)&&(cur!=L.Back()); j++ {
					cur = cur.Next()		
				}
			} else if d < 0 {
				d = -d
				for j:=0; (j<d)&&(cur!=L.Front()); j++ {
					cur = cur.Prev()
				}
			}
		case 2 :
			if cur!=L.Back() {
				cur = cur.Next()
			} else {
				cur = L.Back()
			}
			_ = L.Remove(cur.Prev())
		}
		//for e:=L.Front(); e!=nil; e=e.Next() {
		//	fmt.Print(e.Value," ")
		//}
		//fmt.Println("  cur:",cur.Value)
	}

	for cur=L.Front(); cur!=L.Back(); cur=cur.Next() {
		fmt.Println(cur.Value)
	}
}

