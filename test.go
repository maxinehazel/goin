package main

import (
	"fmt"

	"github.com/maxinekrebs/goin"
)

func main() {
	fileName := "/tmp/example.txt"
	input, err := goin.ReadFromFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(input)
}
