package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {

	file, _ := os.Open(os.Args[1])
	in := bufio.NewReader(file)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("read failed: %v", err)
		}
		log.Printf("%c\n", r)
	}
}
