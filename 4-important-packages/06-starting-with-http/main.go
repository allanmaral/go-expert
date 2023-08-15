package main

import "net/http"

func main() {
	http.HandleFunc("/", FindCep)
	http.ListenAndServe(":8080", nil)
}

func FindCep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
