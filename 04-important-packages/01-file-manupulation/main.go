package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	f.Close()

	fmt.Printf("File created successfully. Size: %d bytes\n", size)

	// file read
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	// file streaming
	file2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 10)
	for {
		n, err2 := reader.Read(buffer)
		if err2 != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// delete file
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
