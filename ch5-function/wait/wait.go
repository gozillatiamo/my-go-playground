package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// os.Exit(1)
		log.Fatalf("Site is down: %v\n", err)
	}
}
