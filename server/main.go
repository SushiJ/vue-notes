package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP port")

	mux := http.NewServeMux()
	mux.HandleFunc("/check", checkHealth)

	// View Notes
	mux.HandleFunc("/notes/view", viewNotes)
	// Create notes
	mux.HandleFunc("/notes/create", createNotes)

	log.Printf("Live at http://localhost%s", *addr)
	log.Fatal((http.ListenAndServe(*addr, mux)))
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200"))
}

func viewNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Notes"))
}

func createNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("Create Notes"))
}
