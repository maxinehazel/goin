package goin

import (
	"fmt"

	"github.com/maxinekrebs/goin"
)

func Example() {
	fileName := "/tmp/example.txt"
	input, err := goin.ReadFromFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(input)
}
