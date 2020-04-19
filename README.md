# Goin

Small go package for reading user input from an editor. Useful for building
CLIs and the like.

Currently supports vim but I am going to expand it to use a generic text editor
eventually.

## Installation

```shell
$ go get github.com/maxinekrebs/goin
```

## Usage

```go
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

```
