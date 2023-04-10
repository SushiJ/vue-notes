package main

import "net/http"

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
