package main

import "fmt"

func main() {
	ch := make(chan string)
	go send("Hello", ch)
	receive(ch)
}

func send(name string, ch chan<- string) {
	// Trying to receive would cause an error
	// <- ch // invalid operation: cannot receive from send-only channel ch
	ch <- name
}

func receive(ch <-chan string) {
	// Trying to send would cause an error
	// ch <- "new message" // invalid operation: cannot send to receive-only channel ch
	fmt.Println(<-ch)
}
