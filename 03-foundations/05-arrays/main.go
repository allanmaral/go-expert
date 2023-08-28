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
	var myArray [3]int
	myArray[0] = 17
	myArray[1] = 42
	myArray[2] = 11

	fmt.Println(myArray[len(myArray)-1])

	for i, v := range myArray {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
