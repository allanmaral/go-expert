package main

import "fmt"

var a int = 0

func main() {
	a := 10

	fmt.Println(a)

	var pointer *int = &a
	*pointer = 20

	fmt.Println(a)

	b := &a

	a = 30

	fmt.Println(*b)
}
