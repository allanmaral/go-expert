package main

// Deadlock example
func main() {
	// Create an empty channel
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}

		// This would be enough
		// forever <- true
	}()

	// Wait for the channel to "fill" resulting in a deadlock
	<-forever
}
