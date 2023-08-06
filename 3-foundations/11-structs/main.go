package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	wesley := Client{
		Name:   "Wesley",
		Age:    30,
		Active: true,
	}
	fmt.Printf("Name: %s, Age: %d, Active: %t\n", wesley.Name, wesley.Age, wesley.Active)

	wesley.Active = false

	fmt.Println(wesley)
}
