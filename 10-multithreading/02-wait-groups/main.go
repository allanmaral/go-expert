package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go task("A", &wg)
	go task("B", &wg)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
}
