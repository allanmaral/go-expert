package main

import (
	"fmt"

	"github.com/allanmaral/go-expert/09-apis/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}
