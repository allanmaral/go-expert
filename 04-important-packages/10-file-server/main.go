package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello from blog"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
