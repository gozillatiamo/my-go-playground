package main

import (
	"encoding/json"
	"fmt"
)

type IA interface {
	GetName() string
}

type A struct {
	name string
	age  int
}

func (A) GetName() string {
	return "name"
}

func (a A) birthYear() int {
	fmt.Println("A func")
	return 2024 - a.age
}

func (a A) backToAge(year int) int {
	return 2024 - year
}

type B struct {
	A
	address string
}

// func (b B) birthYear() int {
// 	fmt.Println("B func")
// 	return 2024 - b.age
// }

func printName(ia IA) {
	fmt.Println(ia.GetName())
}

func main() {

	a := A{}

	printName(a)

	// a := A{"got", 32}
	// fmt.Println(a.birthYear())
	// fmt.Println(a.backToAge(a.birthYear()))
	// b := B{a, "samut sakhon"}
	// // backToAge := A.backToAge
	// // backToAge(b.A, b.birthYear())
	// fmt.Println(b.birthYear())
	// fmt.Println(b.backToAge(b.birthYear()))

	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	var b = make([]int, 3, 7)
	copy(b, a[1:4])
	fmt.Println(len(b), cap(b))
	fmt.Println("Before", a, b)
	b[0] = 20
	fmt.Println("After", a, b)
	type try struct {
		C int `json:"c"`
		B int `json:"b"`
		A int `json:"a"`
	}
	c := map[string]int{"c": 1, "b": 2, "a": 3}
	g := c["a"]
	fmt.Println(g)
	g = 100
	fmt.Println(g)

	fmt.Println(c)
	bStr, _ := json.Marshal(c)
	fmt.Println(string(bStr))

	f := try{1, 2, 3}
	fStr, _ := json.Marshal(f)
	fmt.Println(string(fStr))

	h := [5]int{1, 2, 3, 4, 5}
	z := []*int{}

	for i, v := range h {
		// nV := v
		fmt.Printf("Address value: %p\n", &v)
		z = append(z, &v)
		v += i
	}

	// fmt.Println(h, z)
	x := []int{}
	for _, v := range z {
		x = append(x, *v)
	}
	fmt.Println(h, x)
	// Mac:
	// h = [1,2,3,4,5] mac
	// x = [1,3,5,7,9] mac
	// Tar:
	// h = [1,2,3,4,5]
	// x = [1,3,5,7,9]
	// Tom:
	// h = [1,2,3,4,5]
	// x = [1,2,3,4,5]
}
