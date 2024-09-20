package main

import "fmt"

var tmpDirs = []string{
	"A",
	"B",
	"C",
	"D",
}

func main() {
	var rmdirs []func()
	var values []string
	for _, dir := range tmpDirs {
		// dir := dir
		fmt.Printf("create dir:\t%s\n", dir)
		rmdirs = append(rmdirs, func() {
			fmt.Printf("remove dir:\t%s\n", dir)
		})
		values = append(values, dir)
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
	fmt.Printf("values dir:\t%v\n", values)
}
