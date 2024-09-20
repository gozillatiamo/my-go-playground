package main

import (
	"fmt"
)

type Path struct {
	Points []Point
}

func (p *Path) printNon() {
	p.Points = append(p.Points, Point{3, 3})
	fmt.Println("in non-pointer method: ", p)
}

func (p *Path) printPointer() {
	p.Points = append(p.Points, Point{4, 4})
	fmt.Println("in pointer method: ", p)
}

// func main() {
// 	p := Point{1, 2}
// 	q := Point{2, 6}

// 	ph := Path{
// 		[]Point{p, q},
// 	}
// 	ph.printNon()
// 	fmt.Println("outer method: ", ph)
// 	ph.printPointer()
// 	fmt.Println("outer method: ", ph)
// 	ph.printNon()

// 	// fmt.Println(p.Distance(q))
// }
