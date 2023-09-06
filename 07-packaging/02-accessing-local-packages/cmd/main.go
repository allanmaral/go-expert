package main

import (
	"fmt"

	// To import local packages, we use the package base path and append
	// the local path to the end
	"github.com/allanmaral/go-expert/07-packaging/02-accessing-local-packages/math"
)

func main() {
	m := math.Math{A: 19, B: 23}
	fmt.Println(m.Add())
}
