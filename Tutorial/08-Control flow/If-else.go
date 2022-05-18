package main

import "fmt"

func main() {
	x := 1
	for {
		x++
		if x > 10 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Print(x, " ")
	}
	fmt.Println()
}
