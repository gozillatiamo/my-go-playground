package main

import (
	"fmt"
)

func main() {
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
