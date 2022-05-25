package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		c <- 44
	}()
	fmt.Println(<-c, <-c)

	b := make(chan int, 2)

	b <- 45
	b <- 46

	fmt.Println(<-b, <-b)

	a := make(chan int, 2)

	a <- 45
	a <- 46

	fmt.Println(<-a, <-a)
}
