package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		// print "Hello, World!" message
		fmt.Println("Hello, World!")

		// sleep for 2 seconds
		time.Sleep(2 * time.Second)

		// print "Hello again!" message
		fmt.Println("Hello again!")

		// sleep for 2 seconds
		time.Sleep(2 * time.Second)
	}
}
