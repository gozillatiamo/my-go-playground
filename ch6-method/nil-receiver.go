package main

// An IntList is a linked list of integers.
// A nill *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}
