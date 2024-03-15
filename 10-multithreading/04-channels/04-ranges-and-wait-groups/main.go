package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go publisher(ch, &wg)
	go reader(ch, &wg)

	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

func publisher(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		ch <- i
	}
	wg.Done()
}
