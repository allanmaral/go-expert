package main

import (
	"fmt"
)

func main() {
	value := sum(50, 10)
	fmt.Println(value)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
