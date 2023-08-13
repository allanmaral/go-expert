package main

import (
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(data))
}
