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

func (c Client) Disable() {
	c.Active = false
	fmt.Printf("The client %s was disabled", c.Name)
}

func main() {
	wesley := Client{
		Name:   "Wesley",
		Age:    30,
		Active: true,
	}

	wesley.Disable()
}
