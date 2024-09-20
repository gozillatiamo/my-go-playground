package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer trace("bigSlowOperation")()
	resp := double(2)
	log.Println(resp)
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) (result int) {
	defer func() {
		fmt.Println(result)
		result = 10
	}()
	return x + x
}
