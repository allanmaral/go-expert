package main

import "fmt"

type Address struct {
	Street string
	Number string
	City   string
	State  string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	wesley := Client{
		Name:   "Wesley",
		Age:    30,
		Active: true,
	}
	fmt.Printf("Name: %s, Age: %d, Active: %t\n", wesley.Name, wesley.Age, wesley.Active)

	wesley.Active = false
	wesley.Street = "Street 1"

	fmt.Println(wesley)
}
