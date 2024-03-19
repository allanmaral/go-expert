package main

import (
	"fmt"
	"time"
)

func worker(wid int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", wid, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	workerCount := 10

	for i := 0; i < workerCount; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}
