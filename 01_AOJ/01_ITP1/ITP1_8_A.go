package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	line := bufio.NewScanner(os.Stdin)
	line.Scan()
	str := line.Text() 
	rstr := []rune(str)

	for _, char := range rstr {
		//fmt.Println(string(char))
		if char>=65 && char<=90 {
			fmt.Print(strings.ToLower(string(char)))
		} else if char>=97 && char<=122 {
			fmt.Print(strings.ToUpper(string(char)))
		} else {
			fmt.Print(string(char))
		}
	}
	fmt.Println()
}