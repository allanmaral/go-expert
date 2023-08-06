package main

import "fmt"

type ID int

var (
	a ID      = 10
	b bool    = true
	c int     = 10
	d string  = "John"
	e float64 = 1.2
)

func main() {
	fmt.Printf("A type is %T\n", a)
	fmt.Printf("B type is %T\n", b)
	fmt.Printf("C type is %T\n", c)
	fmt.Printf("D type is %T\n", d)
	fmt.Printf("E type is %T\n", e)
}
