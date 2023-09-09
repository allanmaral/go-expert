package main

import (
	"fmt"

	// We can import local packages telling the go mod to replace this path with a local one
	// $ go mod edit -replace github.com/allanmaral/go-expert/07-packaging/04-working-with-go-mod-replace/math=../math
	"github.com/allanmaral/go-expert/07-packaging/04-working-with-go-mod-replace/math"
)

func main() {
	m := math.New(1, 2)

	fmt.Println(m.Add())
}
