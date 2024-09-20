package main

import "fmt"

type Square struct {
	W, H int
}

func (s *Square) ScaleBy(i int) {
	s.W *= i
	s.H *= i
}

type Box []Square

type Box2 *string

type NString string
type PString *string

func (n *Box) testPrint() {
	fmt.Println(n)
}

func main() {
	s := Square{1, 1}
	s.ScaleBy(5)
	fmt.Printf("Square values is: %#v\n", s)
}
