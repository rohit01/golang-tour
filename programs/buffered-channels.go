package main

import "fmt"

func main() {
	// If buffersize is changed to 1, we get deadlock error
	// fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// if one more print statement is written, we get another deadlock!
}
