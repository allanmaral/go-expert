package main

import "fmt"

type Customer struct {
	name string
}

func (c Customer) walk() {
	fmt.Printf("The customer %s walked", c.name)
}

func main() {
	customer := Customer{
		name: "Jonh",
	}

	customer.walk()
}
