package main

import "fmt"

func main() {
	wages := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(wages["Wesley"])
	delete(wages, "Wesley")
	wages["Wes"] = 5000
	fmt.Println(wages["Wes"])

	sal := make(map[string]int)
	fmt.Println(sal["Allan"])

	for name, wage := range wages {
		fmt.Printf("%s's wage is %d\n", name, wage)
	}

	// blank identifier
	for _, wage := range wages {
		fmt.Printf("wage is %d\n", wage)
	}
}
