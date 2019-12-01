package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

func main() {
	var nextTicket uint64

	http.HandleFunc("/ticket", func(w http.ResponseWriter, r *http.Request) {
		ticket := atomic.AddUint64(&nextTicket, 1)
		ticket--
		fmt.Fprintf(w, "%d", ticket)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
