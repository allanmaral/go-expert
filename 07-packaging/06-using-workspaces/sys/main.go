package main

import (
	"fmt"

	// Using workspaces we can tell the go cli what projects to include, so we dont need
	// to replace in the go.mod
	// $ go work init ./math ./sys
	"github.com/allanmaral/go-expert/07-packaging/04-working-with-go-mod-replace/math"

	"github.com/google/uuid"
)

func main() {
	m := math.New(1, 2)

	fmt.Println(m.Add())

	fmt.Println(uuid.New().String())
}
