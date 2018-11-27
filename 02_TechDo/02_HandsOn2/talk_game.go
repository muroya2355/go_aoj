package main

func main() {

	ch := make(chan string, 0)

	go func() {
		ch <- "name"
		close(ch)
	}()

	theme, closed := <-ch
	fmt.Println()
}
