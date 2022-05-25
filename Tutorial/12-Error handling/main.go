package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	c, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	var str string
	fmt.Print("Enter your text : ")
	fmt.Scan(&str)

	r := strings.NewReader(str)
	io.Copy(c, r)

	_, er := os.Open("no-file.txt")

	if er != nil {
		log.Fatalln(er)
	}
}
