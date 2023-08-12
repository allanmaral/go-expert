package main

import "fmt"

func sum(a, b *int) int {
	*a = 50
	*b = 51
	return *a + *b
}

func main() {
	var1 := 10
	var2 := 20

	fmt.Println(sum(&var1, &var2))
	fmt.Println(var1)
}
