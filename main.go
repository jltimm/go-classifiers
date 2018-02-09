package main

import (
	"fmt"

	"github.com/jltimm/go-classifiers/classifiers"
)

func main() {
	fmt.Printf("Hello, world!\n")
	abs := classifiers.Test()
	fmt.Println(abs)
}
