package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path struct {
	Points []Point
}

// func (path Path) Distance() float64 {
// 	sum := 0.0
// 	for i := range path {
// 		if i > 0 {
// 			sum += path[i-1].Distance(path[i])
// 		}
// 	}
// 	return sum
// }

func (p *Path) printNonP() {
	time.Sleep(2 * time.Second)
	// p.Points[1] = Point{2, 2}
	p.Points = append(p.Points, Point{3, 3})
	fmt.Println("nonP: ", p)
}

func (p *Path) printP() {
	p.Points[0] = Point{1, 1}
	fmt.Println("P: ", p)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	ph := Path{[]Point{p, q}}
	// fmt.Println(p.Distance(q))
	// fmt.Println(Distance(p, q))
	// fmt.Println(ph.Distance())
	// p.ScaleBy(2)
	// fmt.Println(p)
	go ph.printNonP()
	fmt.Println("Original: ", ph)

	ph.printP()
	fmt.Println("Original: ", ph)

	// ph.printNonP()
	// fmt.Println("Original: ", ph)
	time.Sleep(3 * time.Second)
	// fmt.Println("Original: ", ph)
	// ph.printNonP()
	// fmt.Println(ph)
}
