package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

var number uint64 = 0

// run with
// go run -race main.go
// to detect rece condition
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes", number)))
	})
	http.ListenAndServe(":3000", nil)
}

func solutionWithMutex() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes", number)))
	})
	http.ListenAndServe(":3000", nil)
}

func solutionWithAtomicAdd() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essa página %d vezes", number)))
	})
	http.ListenAndServe(":3000", nil)
}
