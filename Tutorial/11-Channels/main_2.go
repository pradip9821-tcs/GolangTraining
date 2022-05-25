package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("code done")
}

func foo(c chan<- int) {
	n := 42
	for i := 0; i < 10; i++ {
		c <- n
		n++
	}
	close(c)
}
