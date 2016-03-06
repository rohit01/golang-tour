package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	go sum([]int{1, 2, 3, 4}, c)
	x := <-c
	y := <-c
	z := <-c
	//x, y := <-c, <-c

	fmt.Println(x, y, z, x+y+z)
}
