package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

var numOfCars = 2

type Car struct {
	name, model, color string
}

var cars = []Car{
	Car{"Toyota", "Corlolla", "red"},
	Car{"Toyota", "Innova", "gray"},
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance
	fmt.Printf("%T\n", distance)

	fmt.Println(distance(p, q))
	a := "interviewbit"
	b := `Interviewbit
	Website`
	fmt.Println(a)
	fmt.Println(b)

	for i := 0; i < numOfCars; i++ {
		numOfCars := 0
		if cars[i].color == "red" {
			numOfCars += 1
			fmt.Println("Inside countRedCars method ", numOfCars)
		}
	}

	arr1 := [2]int{2, 3}
	arr2 := [...]int{2, 3}
	fmt.Println(arr1, arr2)
	fmt.Println(arr1 == arr2)
	fmt.Println(cap(arr1), cap(arr2))

	var x int
	arr := [3]int{3, 5, 2}
	x, arr = arr[0], arr[1:2]
	fmt.Println(x, arr)
}
