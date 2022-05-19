package main

import "fmt"

func main() {
	switch {
	case (1 == 1):
		fmt.Println("hii")
		fallthrough
	case (2 == 3):
		fmt.Println("hii again")
	}

}
