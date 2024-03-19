package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)
	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			id := atomic.AddInt64(&i, 1)
			msg := Message{id, "Hello from RabbitMQ"}
			ch1 <- msg
			time.Sleep(time.Second * 4)
		}
	}()

	// Kafka
	go func() {
		for {
			id := atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 4)
			ch2 <- Message{id, "Hello from Kafka"}
		}
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Printf("Received \"%d - %s\" from channel 1\n", msg.Id, msg.Msg)

		case msg := <-ch2:
			fmt.Printf("Received \"%d - %s\" from channel 2\n", msg.Id, msg.Msg)

		case <-time.After(time.Second * 3):
			fmt.Println("Timeout")

			//default:
			//	fmt.Println("Default")
		}
	}
}
