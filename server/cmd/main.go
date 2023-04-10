package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/check", checkHealth)

	// View Notes
	mux.HandleFunc("/notes/view", viewNotes)
	// Create notes
	mux.HandleFunc("/notes/create", createNotes)

	log.Printf("Live at http://localhost%s", *addr)
	log.Fatal((http.ListenAndServe(*addr, mux)))
}
