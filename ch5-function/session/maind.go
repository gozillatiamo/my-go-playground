package main

import "fmt"

func main() {
	list := []int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}
	sum(list...)
}

func sum2(in []int) int {
	var total int
	for _, i := range in {
		total += i
	}

	fmt.Println(total)
	return total
}

func sum(in ...int) int {
	var total int
	for _, i := range in {
		total += i
	}

	fmt.Println(total)
	return total
}
