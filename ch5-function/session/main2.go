package main

import (
	"errors"
	"fmt"
)

var EOF error = errors.New("EOF")

func main() {
	a := "World"
	Hello(&a)
	fmt.Println(a)
}

func raiseError() error {
	return EOF
}
