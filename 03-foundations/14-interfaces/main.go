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

type Business struct {
}

type Deactivatable interface {
	Disable()
}

func (c Client) Disable() {
	c.Active = false
	fmt.Printf("The client %s was disabled", c.Name)
}

func (b Business) Disable() {
}

func Disabling(d Deactivatable) {
	d.Disable()
}

func main() {
	wesley := Client{
		Name:   "Wesley",
		Age:    30,
		Active: true,
	}

	business := Business{}

	Disabling(wesley)
	Disabling(business)
}
