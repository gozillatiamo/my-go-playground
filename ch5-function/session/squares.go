package main

import "fmt"

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())

	g := squares()
	fmt.Println(g())
	fmt.Println(g())
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
