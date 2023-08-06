package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("the sum is grater then 50")
	}
	return a + b, nil
}
