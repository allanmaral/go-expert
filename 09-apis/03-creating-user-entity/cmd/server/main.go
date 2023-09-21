package main

import (
	"fmt"

	"github.com/allanmaral/go-expert/09-apis/02-creating-configuration-file/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(config)
}
