package main

import "fmt"

func main() {
	ch := make(chan string)

	// New thread
	go func() {
		ch <- "Hello World!"
	}()

	// Main thread
	msg := <-ch
	fmt.Println(msg)
}
