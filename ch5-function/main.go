package main

import (
	"fmt"
	"reflect"
)

// func Sin(x float32) float64 // implemented in assembly language

func fnA(a string, b string) (respA string, respB string) { return }
func fnB(a, b string, c int) (respA, respB string)        { return }

func main() {
	fmt.Printf("Same func:\t%t\n", reflect.TypeOf(fnA) == reflect.TypeOf(fnB))
	sA := map[string]int{
		"A": 1,
		"B": 2,
	}
	A(sA)
	fmt.Printf("%#v\n", sA)
}

func A(a map[string]int) {
	// a[0] = "C"
	// a = append(a, "C")
	a["C"] = 3
}
