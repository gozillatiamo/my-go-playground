package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	// var y uint8 = 1 << 5
	fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(x | y)
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
	fmt.Println("MaxFloat32: ", math.MaxFloat32)
	fmt.Println("MaxFloat64: ", math.MaxFloat64)

	// const e = 2.71828
	fmt.Printf("x = %d e exponent x = %8.3f\n", 5, math.Exp(float64(5)))
	// fmt.Printf("%08b\n", y)
	// fmt.Printf("%08b\n", x|y)

	var f float32 = 16777216

	fmt.Printf("%g == %g: %t\n", f, f-1, f == f-1)
	fmt.Println("Hello, World\rtest")
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	if s == nil {
		fmt.Println("a nil")
	}
	fmt.Printf("Array a:%v\n", a)
	fmt.Printf("Slice s:%v\tlength: %d\tcap: %d\n", s, len(s), cap(s))
	fmt.Println("Change slice value index 1 to 8")
	s[1] = 8
	fmt.Printf("Array a:%v\n", a)
	fmt.Printf("Slice s:%v\tlength: %d\tcap: %d\n", s, len(s), cap(s))
	fmt.Println("Append 7 to slice")
	s = append(s, 7)
	fmt.Printf("Array a:%v\n", a)
	fmt.Printf("Slice s:%v\tlength: %d\tcap: %d\n", s, len(s), cap(s))
	fmt.Println("Append 6 to slice")
	s = append(s, 6)
	fmt.Printf("Array a:%v\n", a)
	fmt.Printf("Slice s:%v\tlength: %d\tcap: %d\n", s, len(s), cap(s))

	fmt.Println("Change slice value index 1 to 0")
	s[1] = 0
	fmt.Printf("Array a:%v\n", a)
	fmt.Printf("Slice s:%v\tlength: %d\tcap: %d\n", s, len(s), cap(s))

}
