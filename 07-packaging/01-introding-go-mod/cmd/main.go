package main

import "fmt"

// To create a package run the command:
// $ go mod init <package path>
//      eg: go mod init github.com/allanmaral/go-expert/07-packaging/01-introducing-go-mod
//
// By convention, the package path match the url where the package is hosted

// We tend to create a `cmd` directory and put the `package main` with the
// entry point there

func main() {
	fmt.Println("Hello, world!")
}
