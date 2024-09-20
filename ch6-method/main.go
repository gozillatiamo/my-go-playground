package main

import (
	"fmt"
)

func main() {
	p := Point{1, 2}
	q := Point{2, 6}

	ph := Path{
		[]Point{p, q},
	}
	ph.printNon()
	fmt.Println("outer method: ", ph)
	ph.printPointer()
	fmt.Println("outer method: ", ph)
	ph.printNon()

	// fmt.Println(p.Distance(q))
}
