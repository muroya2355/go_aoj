package main

import(
	"fmt"
	//"bufio"
	//"os"
	"strings"
)

func main() {

	//line := bufio.NewScanner(os.Stdin)
	alplist := make(map[string]int)
	for i:=int(rune('a')); i<=int(rune('z')); i++ {
		alplist[string(rune(i))] = 0
	}

	var str string
	for {
		//str := line.Text()
		_, err := fmt.Scan(&str)
		if err != nil {
			break
		}
		str = strings.ToLower(str)
		rstr := []rune(str)

		for _, char := range rstr {
			//fmt.Println(string(char))
			alplist[string(char)] += 1
		}
	}
	
	for i:=int(rune('a')); i<=int(rune('z')); i++ {
		fmt.Println(string(rune(i)),":",alplist[string(rune(i))])
	}
}