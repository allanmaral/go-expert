package main

import (
	"fmt"

	// To import local packages, we use the package base path and append
	// the local path to the end
	"github.com/allanmaral/go-expert/07-packaging/03-exporting-objects/math"
)

func main() {
	// Math `a` and `b` are not visible outside the `math` package
	// we need an contructor function to set those values
	m := math.New(19, 23)
	fmt.Println(m.Add())

	// Even the math struct being lower case, we can still acess
	// its public properties
	fmt.Println(m.SomePublic)
}
