package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	for i := 0; i < 5000; i++ {
		f, err := os.Create(path.Join("tmp", fmt.Sprintf("file-%d.txt", i)))
		if err != nil {
			panic(err)
		}

		f.WriteString("hello, world!")
		f.Close()
	}
}
